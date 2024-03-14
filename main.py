from PyQt6.QtWidgets import *
from PyQt6.QtCore import Qt

app = QApplication([])

window = QMainWindow()

# 设置窗口大小
window.setGeometry(100, 100, 1200, 600)

# 创建主垂直布局
main_layout = QVBoxLayout()
main_layout.setSpacing(0)
# 创建表格
table = QTableWidget()

# 设置表格列数
table.setColumnCount(6)

# 设置表格水平表头
table.setHorizontalHeaderLabels(["id", "名称", "类别", "型号", "丝印", "数量"])

# 创建模拟数据
fake_data = [
    ["1", "cpu", "cpu", "411231331313131313131321", "111", "1"],
    ["2", "ram", "ram", "intel", "222", "2"],
    ["33", "ssd", "ssd", "intel", "333", "3"],
    ["44", "hdd", "hdd", "intel", "444", "4"],
]

# 设置表格行数
table.setRowCount(len(fake_data) + 10)

# 添加数据
for i, row in enumerate(fake_data):
    for j, value in enumerate(row):
        item = QTableWidgetItem()

        # 如果是最后一行，将这格设置成可调整大小
        if i == len(fake_data) - 1:
            item.setFlags(item.flags() | Qt.ItemFlag.ItemIsEditable)
            
        item.setText(value)
        table.setItem(i, j, item)

# 控制某一列的宽度
table.setColumnWidth(3, 365)
table.setColumnWidth(4, 365)

# main_layout.setAlignment(Qt.AlignmentFlag.AlignJustify)
# 将表格添加到主垂直布局中
main_layout.addWidget(table)

# 创建水平布局
h_layout = QHBoxLayout()
h_layout.setSpacing(0)
h_layout.setContentsMargins(0, 0, 0, 0)
h_layout.setStretch(0,0)
# 创建翻页按钮
# 创建按钮1
button1 = QPushButton()
button1.setFixedSize( 50, 30)
button1.setText("上一页")

# 创建按钮2
button2 = QPushButton("下一页")
button2.setFixedSize( 50, 30)


# 将按钮添加到水平布局中
h_layout.addWidget(button1)
h_layout.addWidget(button2)

# 将水平布局添加到主垂直布局中
main_layout.addLayout(h_layout)

# 设置窗口中央部件
window.setCentralWidget(QWidget(window))

# 将主垂直布局设置到中央部件
window.centralWidget().setLayout(main_layout)

# 显示窗口
window.show()

# 进入事件循环
app.exec()
