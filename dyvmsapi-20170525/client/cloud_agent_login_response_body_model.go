// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudAgentLoginResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudAgentLoginResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudAgentLoginResponseBody
	GetCode() *string
	SetData(v *CloudAgentLoginResponseBodyData) *CloudAgentLoginResponseBody
	GetData() *CloudAgentLoginResponseBodyData
	SetMessage(v string) *CloudAgentLoginResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudAgentLoginResponseBody
	GetRequestId() *string
}

type CloudAgentLoginResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudAgentLoginResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A90E4451-FED7-49D2-87C8-00700A8C4D0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudAgentLoginResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudAgentLoginResponseBody) GoString() string {
	return s.String()
}

func (s *CloudAgentLoginResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudAgentLoginResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudAgentLoginResponseBody) GetData() *CloudAgentLoginResponseBodyData {
	return s.Data
}

func (s *CloudAgentLoginResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudAgentLoginResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudAgentLoginResponseBody) SetAccessDeniedDetail(v string) *CloudAgentLoginResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudAgentLoginResponseBody) SetCode(v string) *CloudAgentLoginResponseBody {
	s.Code = &v
	return s
}

func (s *CloudAgentLoginResponseBody) SetData(v *CloudAgentLoginResponseBodyData) *CloudAgentLoginResponseBody {
	s.Data = v
	return s
}

func (s *CloudAgentLoginResponseBody) SetMessage(v string) *CloudAgentLoginResponseBody {
	s.Message = &v
	return s
}

func (s *CloudAgentLoginResponseBody) SetRequestId(v string) *CloudAgentLoginResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudAgentLoginResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudAgentLoginResponseBodyData struct {
	// 结果
	//
	// example:
	//
	// 0
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CloudAgentLoginResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudAgentLoginResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudAgentLoginResponseBodyData) GetResult() *int64 {
	return s.Result
}

func (s *CloudAgentLoginResponseBodyData) SetResult(v int64) *CloudAgentLoginResponseBodyData {
	s.Result = &v
	return s
}

func (s *CloudAgentLoginResponseBodyData) Validate() error {
	return dara.Validate(s)
}
