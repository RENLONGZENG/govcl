// 由res2go自动生成，不要编辑。
package main

import (
    "github.com/ying32/govcl/vcl"
)

type TMainForm struct {
    *vcl.TForm
    Splitter1  *vcl.TSplitter
    ToolBar1   *vcl.TToolBar
    StatusBar1 *vcl.TStatusBar
    Panel1     *vcl.TPanel
    Panel3     *vcl.TPanel
    Panel4     *vcl.TPanel `events:"OnPanel3MouseEnter,OnPanel3MouseLeave"`
    Panel5     *vcl.TPanel `events:"OnPanel3MouseEnter,OnPanel3MouseLeave"`
    Panel2     *vcl.TPanel
    Label1     *vcl.TLabel
    Button1    *vcl.TButton
    Button2    *vcl.TButton
    Button3    *vcl.TButton `events:"OnButton2Click"`
    Button4    *vcl.TButton `events:"OnButton2Click"`

    //::private::
    TMainFormFields
}

var MainForm *TMainForm




// 以字节形式加载
// vcl.Application.CreateForm(mainFormBytes, &MainForm)

var (
    mainFormBytes = []byte {
        0x47, 0x4F, 0x56, 0x43, 0x4C, 0x46, 0x4F, 0x52, 0x4D, 0x5A, 0x01, 0x00, 
        0x93, 0x48, 0x3E, 0x1C, 0xDA, 0x61, 0xEB, 0x6B, 0x91, 0xCE, 0x3C, 0x48, 
        0x95, 0x78, 0x9B, 0x9A, 0xE5, 0x9C, 0xB3, 0x5D, 0x7A, 0x04, 0x3F, 0x83, 
        0x2E, 0x3C, 0xEB, 0x7B, 0xFC, 0x4B, 0x2B, 0x54, 0x31, 0xBE, 0xA2, 0x52, 
        0xB0, 0xCE, 0x75, 0x17, 0x9A, 0x75, 0x9C, 0x79, 0xE5, 0xF8, 0xA7, 0x9B, 
        0x13, 0xD5, 0x0F, 0xB1, 0x9C, 0xF2, 0xBD, 0x7A, 0xF7, 0xBD, 0x50, 0xD3, 
        0xB8, 0x1A, 0xCA, 0x52, 0xD3, 0x62, 0x40, 0xED, 0x8C, 0xF5, 0xA8, 0xBB, 
        0x1B, 0x60, 0x51, 0x2F, 0x0F, 0xDE, 0x4E, 0x7B, 0x04, 0xBF, 0xAB, 0xA0, 
        0x3B, 0xDB, 0xD8, 0x45, 0xC3, 0xD0, 0xA8, 0x77, 0x76, 0x5F, 0x9C, 0x04, 
        0x3A, 0x11, 0xFC, 0xF1, 0x96, 0x42, 0xF0, 0x9D, 0x8F, 0x10, 0x98, 0x8A, 
        0x2B, 0xF0, 0x99, 0x5C, 0x04, 0xDC, 0x25, 0xD5, 0x76, 0x6A, 0x1F, 0x1A, 
        0x0D, 0x35, 0x43, 0xCD, 0xBC, 0x68, 0x6F, 0x54, 0xC6, 0xC7, 0x93, 0xB8, 
        0x33, 0xF9, 0xC1, 0xA6, 0x06, 0x78, 0xDC, 0x9B, 0xB3, 0x38, 0x72, 0x79, 
        0x22, 0xFC, 0x75, 0xAE, 0x2A, 0x44, 0x65, 0x96, 0x86, 0x2A, 0xA7, 0x15, 
        0xD7, 0xDE, 0xE7, 0x69, 0x8C, 0x50, 0x36, 0x3E, 0xBA, 0x2A, 0x15, 0x15, 
        0xB2, 0x90, 0xEB, 0x8C, 0xB6, 0x40, 0x53, 0x6C, 0x64, 0x85, 0xA9, 0x89, 
        0x4B, 0xB0, 0x81, 0x88, 0x67, 0x8C, 0x70, 0xA6, 0xD0, 0x7A, 0xD2, 0x0B, 
        0x52, 0xCB, 0x27, 0xE9, 0x3C, 0xAC, 0x0A, 0x9C, 0x82, 0x71, 0x6E, 0x90, 
        0x12, 0x67, 0xA1, 0x4D, 0x67, 0x03, 0xF6, 0x08, 0x26, 0x16, 0x87, 0x30, 
        0xFB, 0x6E, 0x6C, 0x95, 0xE9, 0xA3, 0x8D, 0x36, 0x32, 0xC5, 0xC1, 0xC1, 
        0xC5, 0x60, 0x25, 0x77, 0x7E, 0xF1, 0xCB, 0x75, 0x29, 0x60, 0x71, 0x73, 
        0xE4, 0x35, 0xAB, 0x4B, 0xA1, 0x47, 0x4A, 0x86, 0x7F, 0x4A, 0xB7, 0xD7, 
        0x50, 0x9B, 0x9C, 0x49, 0x6E, 0x92, 0x82, 0xA1, 0xC4, 0xBB, 0xF7, 0x2D, 
        0xFA, 0xBF, 0xFB, 0xAA, 0x58, 0x81, 0x61, 0x8C, 0x72, 0xD8, 0xE3, 0xB2, 
        0x1B, 0x88, 0x89, 0xAA, 0x4F, 0x75, 0xD6, 0xE7, 0xC2, 0x03, 0x05, 0x7D, 
        0x8B, 0x98, 0x81, 0x66, 0x05, 0x65, 0xFD, 0x3B, 0xD4, 0x82, 0x48, 0x4C, 
        0x19, 0xFA, 0xB7, 0x2A, 0x1D, 0x36, 0x84, 0xB8, 0x4D, 0x3D, 0x89, 0x7E, 
        0xC5, 0x00, 0xA3, 0xAD, 0x24, 0x8F, 0x94, 0xAC, 0xD9, 0x92, 0xC9, 0x60, 
        0xE2, 0xFA, 0x47, 0x26, 0x48, 0x4D, 0x45, 0xB3, 0xAA, 0xD7, 0x9D, 0xA6, 
        0x6E, 0x37, 0x45, 0x74, 0x56, 0x8F, 0x50, 0x73, 0x06, 0xDD, 0x27, 0xEE, 
        0x26, 0x17, 0xB1, 0x4D, 0x87, 0xF6, 0x91, 0xA6, 0xEB, 0x5A, 0xFE, 0x6F, 
        0x14, 0x72, 0x16, 0x02, 0xEA, 0xC7, 0xAD, 0xED, 0x9A, 0x16, 0xF6, 0x49, 
        0xF7, 0x4F, 0x01, 0x0D, 0xA2, 0xB3, 0x4D, 0xC8, 0x0B, 0xFE, 0xA9, 0x20, 
        0xF5, 0x7C, 0x49, 0x6C, 0x5E, 0x86, 0xE2, 0x4D, 0xAB, 0x41, 0x69, 0x65, 
        0xB4, 0xB6, 0x35, 0x48, 0xD8, 0xB5, 0x12, 0x25, 0x9A, 0xA9, 0x05, 0x1D, 
        0x78, 0x11, 0x17, 0xB1, 0x20, 0x6B, 0xFE, 0xFD, 0x34, 0x41, 0x43, 0xC7, 
        0x83, 0xBA, 0x49, 0x0B, 0xA1, 0x02, 0x10, 0xED, 0x97, 0xCF, 0x00, 0x77, 
        0x4D, 0xCE, 0x01, 0x35, 0x8B, 0xA3, 0x34, 0x0B, 0xA3, 0x16, 0x84, 0x57, 
        0xF3, 0xD2, 0xFC, 0x76, 0xD9, 0xCE, 0xAA, 0x6D, 0x98, 0xB2, 0x06, 0x86, 
        0xEA, 0x61, 0x2B, 0x65, 0xC2, 0xEA, 0x19, 0xEB, 0xBF, 0x2E, 0x7A, 0x02, 
        0x86, 0xE0, 0x5A, 0x6F, 0x4C, 0xB3, 0x4E, 0x6B, 0x43, 0xD7, 0x2D, 0x11, 
        0x59, 0x1E, 0xC4, 0x1F, 0x3E, 0xD2, 0xEA, 0x71, 0x64, 0x28, 0x13, 0x1F, 
        0x5A, 0xFB, 0x9F, 0x21, 0x7E, 0x2B, 0x39, 0x3B, 0x2C, 0xAA, 0xEB, 0xA7, 
        0x8D, 0x16, 0x3D, 0x68, 0x9C, 0x94, 0xF7, 0x1E, 0xE0, 0x9D, 0x15, 0x67, 
        0x44, 0x67, 0xC4, 0xF0, 0x47, 0xFD, 0x8E, 0xC9, 0x9B, 0x24, 0x13, 0x91, 
        0x7D, 0x00, 0x10, 0x82, 0x2E, 0x9C, 0x4A, 0x4B, 0xD8, 0xB7, 0xE8, 0x77, 
        0x74, 0x74, 0xA7, 0x48, 0xC5, 0xCE, 0x99, 0x3C, 0x6E, 0xDB, 0xCD, 0x41, 
        0xC2, 0xA3, 0x11, 0xC6, 0x07, 0x25, 0xA6, 0xA7, 0x7D, 0x32, 0xA5, 0x30, 
        0x16, 0x76, 0x57, 0x2B, 0x3D, 0x62, 0xB1, 0xD8, 0x37, 0x11, 0xC1, 0x64, 
        0x21, 0xEC, 0x84, 0x87, 0x16, 0xF4, 0x51, 0xC3, 0x1E, 0xEA, 0xCC, 0xB1, 
        0xBC, 0xE3, 0x11, 0xFC, 0x52, 0xE7, 0xD4, 0x39, 0x14, 0x96, 0x25, 0xF4, 
        0xBF, 0x3F, 0x03, 0x31, 0x94, 0x13, 0x19, 0x79, 0xF0, 0xF6, 0x09, 0xB4, 
        0x47, 0xF8, 0x0C, 0xC1, 0x15, 0x25, 0x0C, 0x83, 0x7F, 0x09, 0xC6, 0xEE, 
        0x00, 0xC3, 0xE3, 0x79, 0xDC, 0xBF, 0x87, 0x0F, 0xB2, 0xDF, 0x06, 0xE4, 
        0xCF, 0x5F, 0x0C, 0xC6, 0x27, 0x28, 0x9B, 0xE6, 0x5C, 0xBF, 0x61, 0xE1, 
        0x7E, 0xD3, 0x5D, 0x7A, 0x3F, 0x20, 0x12, 0xE3, 0x70, 0xB7, 0x38, 0x3C, 
        0xEC, 0x19, 0x49, 0x5B, 0x79, 0xBF, 0x53, 0x2B, 0x3A, 0x7F}
)
