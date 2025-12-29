// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOrUpdateWelcomeTextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExtentions(v map[string]interface{}) *AddOrUpdateWelcomeTextResponseBody
	GetExtentions() map[string]interface{}
	SetMessage(v string) *AddOrUpdateWelcomeTextResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddOrUpdateWelcomeTextResponseBody
	GetRequestId() *string
	SetResult(v bool) *AddOrUpdateWelcomeTextResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *AddOrUpdateWelcomeTextResponseBody
	GetStatusCode() *int32
}

type AddOrUpdateWelcomeTextResponseBody struct {
	Extentions map[string]interface{} `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s AddOrUpdateWelcomeTextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddOrUpdateWelcomeTextResponseBody) GoString() string {
	return s.String()
}

func (s *AddOrUpdateWelcomeTextResponseBody) GetExtentions() map[string]interface{} {
	return s.Extentions
}

func (s *AddOrUpdateWelcomeTextResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddOrUpdateWelcomeTextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddOrUpdateWelcomeTextResponseBody) GetResult() *bool {
	return s.Result
}

func (s *AddOrUpdateWelcomeTextResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddOrUpdateWelcomeTextResponseBody) SetExtentions(v map[string]interface{}) *AddOrUpdateWelcomeTextResponseBody {
	s.Extentions = v
	return s
}

func (s *AddOrUpdateWelcomeTextResponseBody) SetMessage(v string) *AddOrUpdateWelcomeTextResponseBody {
	s.Message = &v
	return s
}

func (s *AddOrUpdateWelcomeTextResponseBody) SetRequestId(v string) *AddOrUpdateWelcomeTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddOrUpdateWelcomeTextResponseBody) SetResult(v bool) *AddOrUpdateWelcomeTextResponseBody {
	s.Result = &v
	return s
}

func (s *AddOrUpdateWelcomeTextResponseBody) SetStatusCode(v int32) *AddOrUpdateWelcomeTextResponseBody {
	s.StatusCode = &v
	return s
}

func (s *AddOrUpdateWelcomeTextResponseBody) Validate() error {
	return dara.Validate(s)
}
