// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotlineAgentStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHotlineAgentStatusResponseBody
	GetCode() *string
	SetData(v string) *GetHotlineAgentStatusResponseBody
	GetData() *string
	SetHttpStatusCode(v int64) *GetHotlineAgentStatusResponseBody
	GetHttpStatusCode() *int64
	SetMessage(v string) *GetHotlineAgentStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotlineAgentStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetHotlineAgentStatusResponseBody
	GetSuccess() *bool
}

type GetHotlineAgentStatusResponseBody struct {
	// Status code. A value of "Success" indicates that the request succeeded.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Agent status. Valid values:
	//
	// - **Offline*	-
	//
	// - **Idle*	-
	//
	// - **Post-processing*	-
	//
	// - **In Call*	-
	//
	// - **Break - Rest*	-
	//
	// - **Break - Meal*	-
	//
	// - **Break - Meeting*	-
	//
	// - **Break - Training*	-
	//
	// - **Break - Coaching**
	//
	// example:
	//
	// 通话中
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the API call succeeded. Valid values:
	//
	// - **true**: Succeeded.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHotlineAgentStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotlineAgentStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotlineAgentStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHotlineAgentStatusResponseBody) GetData() *string {
	return s.Data
}

func (s *GetHotlineAgentStatusResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *GetHotlineAgentStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotlineAgentStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotlineAgentStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetHotlineAgentStatusResponseBody) SetCode(v string) *GetHotlineAgentStatusResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotlineAgentStatusResponseBody) SetData(v string) *GetHotlineAgentStatusResponseBody {
	s.Data = &v
	return s
}

func (s *GetHotlineAgentStatusResponseBody) SetHttpStatusCode(v int64) *GetHotlineAgentStatusResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetHotlineAgentStatusResponseBody) SetMessage(v string) *GetHotlineAgentStatusResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotlineAgentStatusResponseBody) SetRequestId(v string) *GetHotlineAgentStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotlineAgentStatusResponseBody) SetSuccess(v bool) *GetHotlineAgentStatusResponseBody {
	s.Success = &v
	return s
}

func (s *GetHotlineAgentStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
