// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRumExceptionStackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetRumExceptionStackResponseBody
	GetCode() *string
	SetData(v *GetRumExceptionStackResponseBodyData) *GetRumExceptionStackResponseBody
	GetData() *GetRumExceptionStackResponseBodyData
	SetHttpStatusCode(v string) *GetRumExceptionStackResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *GetRumExceptionStackResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetRumExceptionStackResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetRumExceptionStackResponseBody
	GetSuccess() *string
}

type GetRumExceptionStackResponseBody struct {
	// The responses code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	Data *GetRumExceptionStackResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// Internal error. Please try again. Contact the DingTalk service account if the issue                              persists after multiple retries.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request.
	//
	// example:
	//
	// B6A00968-82A8-4F14-9D1B-B53827DB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetRumExceptionStackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRumExceptionStackResponseBody) GoString() string {
	return s.String()
}

func (s *GetRumExceptionStackResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetRumExceptionStackResponseBody) GetData() *GetRumExceptionStackResponseBodyData {
	return s.Data
}

func (s *GetRumExceptionStackResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *GetRumExceptionStackResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetRumExceptionStackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRumExceptionStackResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetRumExceptionStackResponseBody) SetCode(v string) *GetRumExceptionStackResponseBody {
	s.Code = &v
	return s
}

func (s *GetRumExceptionStackResponseBody) SetData(v *GetRumExceptionStackResponseBodyData) *GetRumExceptionStackResponseBody {
	s.Data = v
	return s
}

func (s *GetRumExceptionStackResponseBody) SetHttpStatusCode(v string) *GetRumExceptionStackResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetRumExceptionStackResponseBody) SetMessage(v string) *GetRumExceptionStackResponseBody {
	s.Message = &v
	return s
}

func (s *GetRumExceptionStackResponseBody) SetRequestId(v string) *GetRumExceptionStackResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRumExceptionStackResponseBody) SetSuccess(v string) *GetRumExceptionStackResponseBody {
	s.Success = &v
	return s
}

func (s *GetRumExceptionStackResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetRumExceptionStackResponseBodyData struct {
	// The name and UUID of the symbol table required for parsing the exception stack. This parameter is exposed during the parsing of PC errors.
	//
	// example:
	//
	// "04B5B216682E40BF9BBE9698E3F98CAA0,libcurl.4.dylib;7878DB3CF21A3C13A203B7E3B0FA66250,libalibabacloud_rum.dylib;0F9F96FE6B1C3253A33AC9E4A0C2A3860,libsystem_kernel.dylib;3DF3256F466E37BCB995A5A9956E14150,libsystem_pthread.dylib;000000000000000000000000000000000,Security;EA4B83A319EB3E15B22CDF035DBD49250,alibabacloud_rum_example;710BB12EEEC744BAB41D1849CA3AD8021,LTSDK.pdb;EE330BA9C49E4730AA15A2B7C0BB2CAE1,JBLive.pdb"
	BinaryImages *string `json:"BinaryImages,omitempty" xml:"BinaryImages,omitempty"`
	// The crash address. This parameter is exposed during the parsing of PC errors.
	//
	// example:
	//
	// 0x1
	CrashAddress *string `json:"CrashAddress,omitempty" xml:"CrashAddress,omitempty"`
	// The cause of the exception. This parameter is exposed during the parsing of PC errors.
	//
	// example:
	//
	// EXC_BAD_ACCESS / KERN_INVALID_ADDRESS
	CrashReason *string `json:"CrashReason,omitempty" xml:"CrashReason,omitempty"`
	// The list of stacks.
	Lines []*string `json:"Lines,omitempty" xml:"Lines,omitempty" type:"Repeated"`
	// The name of the crash parsing module. This parameter is exposed during the parsing of PC errors.
	//
	// example:
	//
	// alibabacloud_rum_example
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// The thread ID.
	//
	// example:
	//
	// 16643
	ThreadId *string `json:"ThreadId,omitempty" xml:"ThreadId,omitempty"`
	// The thread stack information captured during PC crashes.
	ThreadInfoList []*GetRumExceptionStackResponseBodyDataThreadInfoList `json:"ThreadInfoList,omitempty" xml:"ThreadInfoList,omitempty" type:"Repeated"`
	// The UUID of the symbol table required for parsing the stack. This parameter is exposed during the parsing of PC errors.
	//
	// example:
	//
	// 9032259CEB9130E780C6DE8FDECCD7990
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetRumExceptionStackResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetRumExceptionStackResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetRumExceptionStackResponseBodyData) GetBinaryImages() *string {
	return s.BinaryImages
}

func (s *GetRumExceptionStackResponseBodyData) GetCrashAddress() *string {
	return s.CrashAddress
}

func (s *GetRumExceptionStackResponseBodyData) GetCrashReason() *string {
	return s.CrashReason
}

func (s *GetRumExceptionStackResponseBodyData) GetLines() []*string {
	return s.Lines
}

func (s *GetRumExceptionStackResponseBodyData) GetModuleName() *string {
	return s.ModuleName
}

func (s *GetRumExceptionStackResponseBodyData) GetThreadId() *string {
	return s.ThreadId
}

func (s *GetRumExceptionStackResponseBodyData) GetThreadInfoList() []*GetRumExceptionStackResponseBodyDataThreadInfoList {
	return s.ThreadInfoList
}

func (s *GetRumExceptionStackResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *GetRumExceptionStackResponseBodyData) SetBinaryImages(v string) *GetRumExceptionStackResponseBodyData {
	s.BinaryImages = &v
	return s
}

func (s *GetRumExceptionStackResponseBodyData) SetCrashAddress(v string) *GetRumExceptionStackResponseBodyData {
	s.CrashAddress = &v
	return s
}

func (s *GetRumExceptionStackResponseBodyData) SetCrashReason(v string) *GetRumExceptionStackResponseBodyData {
	s.CrashReason = &v
	return s
}

func (s *GetRumExceptionStackResponseBodyData) SetLines(v []*string) *GetRumExceptionStackResponseBodyData {
	s.Lines = v
	return s
}

func (s *GetRumExceptionStackResponseBodyData) SetModuleName(v string) *GetRumExceptionStackResponseBodyData {
	s.ModuleName = &v
	return s
}

func (s *GetRumExceptionStackResponseBodyData) SetThreadId(v string) *GetRumExceptionStackResponseBodyData {
	s.ThreadId = &v
	return s
}

func (s *GetRumExceptionStackResponseBodyData) SetThreadInfoList(v []*GetRumExceptionStackResponseBodyDataThreadInfoList) *GetRumExceptionStackResponseBodyData {
	s.ThreadInfoList = v
	return s
}

func (s *GetRumExceptionStackResponseBodyData) SetUuid(v string) *GetRumExceptionStackResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *GetRumExceptionStackResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type GetRumExceptionStackResponseBodyDataThreadInfoList struct {
	// Thread stack details.
	//
	// example:
	//
	// "0  libsystem_platform.dylib + 0x1ab5\\n    rax = 0x0000000000000001   rdx = 0x0000000000000064\\n    rcx = 0xffffffffffffffff   rbx = 0x0000000107701bd0\\n    rsi = 0x0101010101010101   rdi = 0x0000000000000001\\n    rbp = 0x00007ff7b8d64300   rsp = 0x00007ff7b8d64300\\n     r8 = 0x000000000000000a    r9 = 0x0000000000000000\\n    r10 = 0x0000000000000001   r11 = 0x0000000000000247\\n    r12 = 0x00007ff7b8d64390   r13 = 0x0000000000000000\\n    r14 = 0x000000010719d770   r15 = 0x00007ff7b8d64500\\n    rip = 0x00007ff807a40ab5\\n    Found by: given as instruction pointer in context\\n 1  alibabacloud_rum_example + 0x2ad1\\n    rbp = 0x00007ff7b8d64310   rsp = 0x00007ff7b8d64310\\n    rip = 0x000000010719dad1\\n    Found by: previous frame\\"s frame pointer\\n 2  alibabacloud_rum_example + 0x2a3b\\n    rbp = 0x00007ff7b8d64360   rsp = 0x00007ff7b8d64320\\n    rip = 0x000000010719da3b\\n    Found by: previous frame\\"s frame pointer\\n 3  0x7ff807688345\\n    rbp = 0x00007ff7b8d64580   rsp = 0x00007ff7b8d64370\\n    rip = 0x00007ff807688345\\n    Found by: previous frame\\"s frame pointer"
	ThreadDetail *string `json:"ThreadDetail,omitempty" xml:"ThreadDetail,omitempty"`
	// The thread tag, including the thread number and name.
	//
	// example:
	//
	// Thread 0 (crashed)
	ThreadTag *string `json:"ThreadTag,omitempty" xml:"ThreadTag,omitempty"`
}

func (s GetRumExceptionStackResponseBodyDataThreadInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetRumExceptionStackResponseBodyDataThreadInfoList) GoString() string {
	return s.String()
}

func (s *GetRumExceptionStackResponseBodyDataThreadInfoList) GetThreadDetail() *string {
	return s.ThreadDetail
}

func (s *GetRumExceptionStackResponseBodyDataThreadInfoList) GetThreadTag() *string {
	return s.ThreadTag
}

func (s *GetRumExceptionStackResponseBodyDataThreadInfoList) SetThreadDetail(v string) *GetRumExceptionStackResponseBodyDataThreadInfoList {
	s.ThreadDetail = &v
	return s
}

func (s *GetRumExceptionStackResponseBodyDataThreadInfoList) SetThreadTag(v string) *GetRumExceptionStackResponseBodyDataThreadInfoList {
	s.ThreadTag = &v
	return s
}

func (s *GetRumExceptionStackResponseBodyDataThreadInfoList) Validate() error {
	return dara.Validate(s)
}
