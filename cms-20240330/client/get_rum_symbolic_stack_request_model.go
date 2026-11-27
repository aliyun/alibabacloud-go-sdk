// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRumSymbolicStackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileInfo(v string) *GetRumSymbolicStackRequest
	GetFileInfo() *string
	SetImages(v string) *GetRumSymbolicStackRequest
	GetImages() *string
	SetPid(v string) *GetRumSymbolicStackRequest
	GetPid() *string
	SetServiceId(v string) *GetRumSymbolicStackRequest
	GetServiceId() *string
	SetStack(v string) *GetRumSymbolicStackRequest
	GetStack() *string
	SetStackType(v string) *GetRumSymbolicStackRequest
	GetStackType() *string
	SetThreadId(v string) *GetRumSymbolicStackRequest
	GetThreadId() *string
	SetWorkspace(v string) *GetRumSymbolicStackRequest
	GetWorkspace() *string
}

type GetRumSymbolicStackRequest struct {
	// The symbol file matching information, which varies based on stackType. For js, jvm, and android-native, it is a JSON string (js and jvm must contain both version and fileName, and optionally uuid/debugId; android-native supports a single module, a modules array, and il2cppMapping). For cocoa, it is a non-JSON format (module:objectKey;...) for historical compatibility only, and it is recommended to pass it through images in the new version. If matching information is missing, the symbol is not found, or parsing fails for js, an empty result is returned. If fileInfo is missing or the symbol is not found for jvm, an empty result is returned, and it falls back to the original stack only when downstream symbolication fails.
	//
	// example:
	//
	// {"version":"1.0.0","fileName":"mapping.txt"}
	FileInfo *string `json:"fileInfo,omitempty" xml:"fileInfo,omitempty"`
	// The binary image information associated with Native/Cocoa crashes. The format varies based on stackType (only electron-native uses a JSON array). For cocoa, it is a line-by-line string (each line is module:arch;uuid;loadAddress). For electron-native, it is a JSON array string (elements contain name, debug_id/debugId, or code_id/codeId, and optionally load_address/loadAddress and arch). For windows-native, it is a semicolon-separated string of debugId and debugFile (.pdb file name). For macos-native, it is a semicolon-separated string of uuid and name. If matching information is missing, the symbol file is not found, or symbolication fails, Native/Cocoa types fall back to returning the original stack lines. The following example applies only to scenarios where stackType is electron-native and cannot be directly combined with stackType=jvm in a request.
	//
	// example:
	//
	// [{"name":"Electron Framework","debug_id":"ABCDEF01-2345-6789-ABCD-EF0123456789","code_id":"CODE-1","load_address":"0x0000000100000000","arch":"arm64"}]
	Images *string `json:"images,omitempty" xml:"images,omitempty"`
	// Deprecated
	//
	// The application PID. **[Deprecated]*	- The server automatically resolves this value based on the workspace and serviceId.
	//
	// example:
	//
	// xxxojw@e3hd38hxxxx11
	Pid *string `json:"pid,omitempty" xml:"pid,omitempty"`
	// The service ID of the RUM application. This parameter is required and cannot be an empty string. It must be used together with workspace to locate the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// a1b2c3d4e5@f67890
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The original exception stack content to be symbolicated. This parameter is required and cannot be an empty string. If stackType is js, the format is "line,column,rows". For other types, pass the original stack text of the corresponding platform.
	//
	// This parameter is required.
	//
	// example:
	//
	// java.lang.NullPointerException
	//
	// 	at com.example.Main.main(Main.java:10)
	Stack *string `json:"stack,omitempty" xml:"stack,omitempty"`
	// The stack type. This parameter is required and cannot be an empty string. Valid values: js, jvm, android-native, cocoa, electron-native, windows-native, and macos-native.
	//
	// This parameter is required.
	//
	// example:
	//
	// jvm
	StackType *string `json:"stackType,omitempty" xml:"stackType,omitempty"`
	// The ID of the exception thread.
	//
	// example:
	//
	// thread-1
	ThreadId *string `json:"threadId,omitempty" xml:"threadId,omitempty"`
	// The name of the workspace to which the RUM application belongs. This parameter is required and cannot be an empty string. It must be used together with serviceId to locate the application.
	//
	// This parameter is required.
	//
	// example:
	//
	// default-cms-1234567890-cn-hangzhou
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetRumSymbolicStackRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRumSymbolicStackRequest) GoString() string {
	return s.String()
}

func (s *GetRumSymbolicStackRequest) GetFileInfo() *string {
	return s.FileInfo
}

func (s *GetRumSymbolicStackRequest) GetImages() *string {
	return s.Images
}

func (s *GetRumSymbolicStackRequest) GetPid() *string {
	return s.Pid
}

func (s *GetRumSymbolicStackRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetRumSymbolicStackRequest) GetStack() *string {
	return s.Stack
}

func (s *GetRumSymbolicStackRequest) GetStackType() *string {
	return s.StackType
}

func (s *GetRumSymbolicStackRequest) GetThreadId() *string {
	return s.ThreadId
}

func (s *GetRumSymbolicStackRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetRumSymbolicStackRequest) SetFileInfo(v string) *GetRumSymbolicStackRequest {
	s.FileInfo = &v
	return s
}

func (s *GetRumSymbolicStackRequest) SetImages(v string) *GetRumSymbolicStackRequest {
	s.Images = &v
	return s
}

func (s *GetRumSymbolicStackRequest) SetPid(v string) *GetRumSymbolicStackRequest {
	s.Pid = &v
	return s
}

func (s *GetRumSymbolicStackRequest) SetServiceId(v string) *GetRumSymbolicStackRequest {
	s.ServiceId = &v
	return s
}

func (s *GetRumSymbolicStackRequest) SetStack(v string) *GetRumSymbolicStackRequest {
	s.Stack = &v
	return s
}

func (s *GetRumSymbolicStackRequest) SetStackType(v string) *GetRumSymbolicStackRequest {
	s.StackType = &v
	return s
}

func (s *GetRumSymbolicStackRequest) SetThreadId(v string) *GetRumSymbolicStackRequest {
	s.ThreadId = &v
	return s
}

func (s *GetRumSymbolicStackRequest) SetWorkspace(v string) *GetRumSymbolicStackRequest {
	s.Workspace = &v
	return s
}

func (s *GetRumSymbolicStackRequest) Validate() error {
	return dara.Validate(s)
}
