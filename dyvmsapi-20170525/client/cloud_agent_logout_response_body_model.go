// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudAgentLogoutResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudAgentLogoutResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudAgentLogoutResponseBody
	GetCode() *string
	SetData(v *CloudAgentLogoutResponseBodyData) *CloudAgentLogoutResponseBody
	GetData() *CloudAgentLogoutResponseBodyData
	SetMessage(v string) *CloudAgentLogoutResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudAgentLogoutResponseBody
	GetRequestId() *string
}

type CloudAgentLogoutResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudAgentLogoutResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7BF47617-7851-48F7-A3A1-2021342A78E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudAgentLogoutResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudAgentLogoutResponseBody) GoString() string {
	return s.String()
}

func (s *CloudAgentLogoutResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudAgentLogoutResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudAgentLogoutResponseBody) GetData() *CloudAgentLogoutResponseBodyData {
	return s.Data
}

func (s *CloudAgentLogoutResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudAgentLogoutResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudAgentLogoutResponseBody) SetAccessDeniedDetail(v string) *CloudAgentLogoutResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudAgentLogoutResponseBody) SetCode(v string) *CloudAgentLogoutResponseBody {
	s.Code = &v
	return s
}

func (s *CloudAgentLogoutResponseBody) SetData(v *CloudAgentLogoutResponseBodyData) *CloudAgentLogoutResponseBody {
	s.Data = v
	return s
}

func (s *CloudAgentLogoutResponseBody) SetMessage(v string) *CloudAgentLogoutResponseBody {
	s.Message = &v
	return s
}

func (s *CloudAgentLogoutResponseBody) SetRequestId(v string) *CloudAgentLogoutResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudAgentLogoutResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudAgentLogoutResponseBodyData struct {
	// 结果
	//
	// example:
	//
	// 0
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CloudAgentLogoutResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudAgentLogoutResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudAgentLogoutResponseBodyData) GetResult() *int64 {
	return s.Result
}

func (s *CloudAgentLogoutResponseBodyData) SetResult(v int64) *CloudAgentLogoutResponseBodyData {
	s.Result = &v
	return s
}

func (s *CloudAgentLogoutResponseBodyData) Validate() error {
	return dara.Validate(s)
}
