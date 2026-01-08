// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetViberPauseTimesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetViberPauseTimesResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetViberPauseTimesResponseBody
	GetCode() *string
	SetData(v map[string]interface{}) *GetViberPauseTimesResponseBody
	GetData() map[string]interface{}
	SetMessage(v string) *GetViberPauseTimesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetViberPauseTimesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetViberPauseTimesResponseBody
	GetSuccess() *bool
}

type GetViberPauseTimesResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 11
	Data map[string]interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 90E63D28-E31D-1EB2**9-A94866411B2O
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetViberPauseTimesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetViberPauseTimesResponseBody) GoString() string {
	return s.String()
}

func (s *GetViberPauseTimesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetViberPauseTimesResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetViberPauseTimesResponseBody) GetData() map[string]interface{} {
	return s.Data
}

func (s *GetViberPauseTimesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetViberPauseTimesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetViberPauseTimesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetViberPauseTimesResponseBody) SetAccessDeniedDetail(v string) *GetViberPauseTimesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetViberPauseTimesResponseBody) SetCode(v string) *GetViberPauseTimesResponseBody {
	s.Code = &v
	return s
}

func (s *GetViberPauseTimesResponseBody) SetData(v map[string]interface{}) *GetViberPauseTimesResponseBody {
	s.Data = v
	return s
}

func (s *GetViberPauseTimesResponseBody) SetMessage(v string) *GetViberPauseTimesResponseBody {
	s.Message = &v
	return s
}

func (s *GetViberPauseTimesResponseBody) SetRequestId(v string) *GetViberPauseTimesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetViberPauseTimesResponseBody) SetSuccess(v bool) *GetViberPauseTimesResponseBody {
	s.Success = &v
	return s
}

func (s *GetViberPauseTimesResponseBody) Validate() error {
	return dara.Validate(s)
}
