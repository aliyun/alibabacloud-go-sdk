// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotspotAnalysisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHotspotAnalysisResponseBody
	GetCode() *string
	SetData(v string) *GetHotspotAnalysisResponseBody
	GetData() *string
	SetMessage(v string) *GetHotspotAnalysisResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotspotAnalysisResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetHotspotAnalysisResponseBody
	GetSuccess() *bool
}

type GetHotspotAnalysisResponseBody struct {
	// error code
	//
	// example:
	//
	// SysomOpenAPI.ServerError
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Returned Data
	//
	// example:
	//
	// "AliYunDun:1657494进程(1657494):根据提供的热点调用栈信息，热点主要集中在`__nanosleep_nocancel`以及相关的系统调用`SyS_nanosleep`上，这表明进程在等待特定的时间间隔。通常这是由于应用中存在长时间的睡眠或高频率的定时操作导致的，比如定时任务、心跳检测等。\\n\\n应用代码在需要定期执行某些检查或等待外部事件时会使用`nanosleep`来实现精确的延时控制。为了优化这种情况，可以考虑以下方案：\\n- 评估是否可以减少定时任务的频率。\\n- 使用条件变量替代单纯的睡眠等待，以响应更快的事件触发。\\n- 如果是I/O密集型操作等待，考虑优化I/O路径或提升I/O效率。建议使用SysOM平台的IO诊断工具来进一步定位具体的I/O瓶颈。"
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// Response message
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2D693121-C925-5154-8DF6-C09A8B369822
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the invocation is successful
	//
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetHotspotAnalysisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotspotAnalysisResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotspotAnalysisResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHotspotAnalysisResponseBody) GetData() *string {
	return s.Data
}

func (s *GetHotspotAnalysisResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotspotAnalysisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotspotAnalysisResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetHotspotAnalysisResponseBody) SetCode(v string) *GetHotspotAnalysisResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotspotAnalysisResponseBody) SetData(v string) *GetHotspotAnalysisResponseBody {
	s.Data = &v
	return s
}

func (s *GetHotspotAnalysisResponseBody) SetMessage(v string) *GetHotspotAnalysisResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotspotAnalysisResponseBody) SetRequestId(v string) *GetHotspotAnalysisResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotspotAnalysisResponseBody) SetSuccess(v bool) *GetHotspotAnalysisResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotspotAnalysisResponseBody) Validate() error {
	return dara.Validate(s)
}
