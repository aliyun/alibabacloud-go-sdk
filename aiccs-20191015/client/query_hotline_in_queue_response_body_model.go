// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHotlineInQueueResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryHotlineInQueueResponseBody
	GetCode() *string
	SetData(v string) *QueryHotlineInQueueResponseBody
	GetData() *string
	SetMessage(v string) *QueryHotlineInQueueResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryHotlineInQueueResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryHotlineInQueueResponseBody
	GetSuccess() *bool
}

type QueryHotlineInQueueResponseBody struct {
	// Status code. A return value of 200 indicates that the request succeeded.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Hotline agent data.
	//
	// example:
	//
	// {"n_online_now":10,"n_idle_now":3,"n_resttype_now":3,"n_acw_now":1}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Status code description.
	//
	// example:
	//
	// xxxx
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// EE338D98-9BD3-4413-B165
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the API call succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryHotlineInQueueResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryHotlineInQueueResponseBody) GoString() string {
	return s.String()
}

func (s *QueryHotlineInQueueResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryHotlineInQueueResponseBody) GetData() *string {
	return s.Data
}

func (s *QueryHotlineInQueueResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryHotlineInQueueResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryHotlineInQueueResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryHotlineInQueueResponseBody) SetCode(v string) *QueryHotlineInQueueResponseBody {
	s.Code = &v
	return s
}

func (s *QueryHotlineInQueueResponseBody) SetData(v string) *QueryHotlineInQueueResponseBody {
	s.Data = &v
	return s
}

func (s *QueryHotlineInQueueResponseBody) SetMessage(v string) *QueryHotlineInQueueResponseBody {
	s.Message = &v
	return s
}

func (s *QueryHotlineInQueueResponseBody) SetRequestId(v string) *QueryHotlineInQueueResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryHotlineInQueueResponseBody) SetSuccess(v bool) *QueryHotlineInQueueResponseBody {
	s.Success = &v
	return s
}

func (s *QueryHotlineInQueueResponseBody) Validate() error {
	return dara.Validate(s)
}
