// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRumExceptionStackRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExceptionBinaryImages(v string) *GetRumExceptionStackRequest
	GetExceptionBinaryImages() *string
	SetExceptionStack(v string) *GetRumExceptionStackRequest
	GetExceptionStack() *string
	SetExceptionThreadId(v string) *GetRumExceptionStackRequest
	GetExceptionThreadId() *string
	SetExtraInfo(v string) *GetRumExceptionStackRequest
	GetExtraInfo() *string
	SetPid(v string) *GetRumExceptionStackRequest
	GetPid() *string
	SetRegionId(v string) *GetRumExceptionStackRequest
	GetRegionId() *string
	SetSourcemapType(v string) *GetRumExceptionStackRequest
	GetSourcemapType() *string
}

type GetRumExceptionStackRequest struct {
	// The binary images, which represent all executable files loaded into the process address space when a crash occurs.
	//
	// example:
	//
	// iOSDemo:arm64%3B1489F4D3-6DE2-300C-90E9-E1B869675351%3B0x0000000104064000\\nAlibabaCloudRUM:arm64%3BAB7B3A8E-6CEE-325D-BCBB-8DA50E61804F%3B0x0000000106660000\\nlibdispatch.dylib:arm
	ExceptionBinaryImages *string `json:"ExceptionBinaryImages,omitempty" xml:"ExceptionBinaryImages,omitempty"`
	// The exception stack information. Set the value to a JSON string. call_stack.info represents the stack information, call_stack.thread.name represents the thread name, and call_stack.thread.id represents the thread ID. This parameter is exactly the same as the exception.stack parameter in the logstore-rum Logstore of Simple Log Service.
	//
	// example:
	//
	// [
	//
	// {
	//
	//     "call_stack.info": "libsystem_kernel.dylib  0x00000001f1ce9178 0x00000001f1ce8000 + 4472\\r\\nlibsystem_kernel.dylib  0x00000001f1ce8f10 0x00000001f1ce8000 + 3856\\r\\nlibsystem_kernel.dylib  0x00000001f1ced718 0x00000001f1ce8000 + 22296\\r\\nAlibabaCloudRUM  0x0000000106711af4 0x0000000106660000 + 727796\\r\\nlibsystem_pthread.dylib  0x00000002146744d4 0x0000000214672000 + 9428",
	//
	//     "call_stack.thread.name": "#3 BRSCrash Exception Handler (Secondary)",
	//
	//     "call_stack.thread.id": "16643"
	//
	//   }
	//
	// ]
	ExceptionStack *string `json:"ExceptionStack,omitempty" xml:"ExceptionStack,omitempty"`
	// The ID of the exception thread.
	//
	// example:
	//
	// 16643
	ExceptionThreadId *string `json:"ExceptionThreadId,omitempty" xml:"ExceptionThreadId,omitempty"`
	// Extra information about iOS symbol tables. You can leave this parameter empty.
	//
	// example:
	//
	// GraphicsServices:system/GraphicsServices/85419099-269B-336D-86B4-0D52D0FF6923/GraphicsServices;WebCore:system/WebCore/BF44A3F4-85D4-38C8-BF26-197F06ADE273/WebCore
	ExtraInfo *string `json:"ExtraInfo,omitempty" xml:"ExtraInfo,omitempty"`
	// The application ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// atxxxxzkcf@d8deedfa9bf****
	Pid *string `json:"Pid,omitempty" xml:"Pid,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The parsing type. Valid values:
	//
	// 	- js: Parses JavaScript errors.
	//
	// 	- sym: Parses PC errors.
	//
	// 	- har: Parses HarmonyOS errors.
	//
	// 	- dSYM: Parses iOS errors.
	//
	// 	- so: Parses Android errors.
	//
	// example:
	//
	// source-map
	SourcemapType *string `json:"SourcemapType,omitempty" xml:"SourcemapType,omitempty"`
}

func (s GetRumExceptionStackRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRumExceptionStackRequest) GoString() string {
	return s.String()
}

func (s *GetRumExceptionStackRequest) GetExceptionBinaryImages() *string {
	return s.ExceptionBinaryImages
}

func (s *GetRumExceptionStackRequest) GetExceptionStack() *string {
	return s.ExceptionStack
}

func (s *GetRumExceptionStackRequest) GetExceptionThreadId() *string {
	return s.ExceptionThreadId
}

func (s *GetRumExceptionStackRequest) GetExtraInfo() *string {
	return s.ExtraInfo
}

func (s *GetRumExceptionStackRequest) GetPid() *string {
	return s.Pid
}

func (s *GetRumExceptionStackRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetRumExceptionStackRequest) GetSourcemapType() *string {
	return s.SourcemapType
}

func (s *GetRumExceptionStackRequest) SetExceptionBinaryImages(v string) *GetRumExceptionStackRequest {
	s.ExceptionBinaryImages = &v
	return s
}

func (s *GetRumExceptionStackRequest) SetExceptionStack(v string) *GetRumExceptionStackRequest {
	s.ExceptionStack = &v
	return s
}

func (s *GetRumExceptionStackRequest) SetExceptionThreadId(v string) *GetRumExceptionStackRequest {
	s.ExceptionThreadId = &v
	return s
}

func (s *GetRumExceptionStackRequest) SetExtraInfo(v string) *GetRumExceptionStackRequest {
	s.ExtraInfo = &v
	return s
}

func (s *GetRumExceptionStackRequest) SetPid(v string) *GetRumExceptionStackRequest {
	s.Pid = &v
	return s
}

func (s *GetRumExceptionStackRequest) SetRegionId(v string) *GetRumExceptionStackRequest {
	s.RegionId = &v
	return s
}

func (s *GetRumExceptionStackRequest) SetSourcemapType(v string) *GetRumExceptionStackRequest {
	s.SourcemapType = &v
	return s
}

func (s *GetRumExceptionStackRequest) Validate() error {
	return dara.Validate(s)
}
