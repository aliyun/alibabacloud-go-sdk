// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkBindClientTelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkBindClientTelResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkBindClientTelResponseBody
	GetCode() *string
	SetData(v *ClinkBindClientTelResponseBodyData) *ClinkBindClientTelResponseBody
	GetData() *ClinkBindClientTelResponseBodyData
	SetMessage(v string) *ClinkBindClientTelResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkBindClientTelResponseBody
	GetRequestId() *string
}

type ClinkBindClientTelResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkBindClientTelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7BF47617-7851-48F7-A3A1-2021342A78E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkBindClientTelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkBindClientTelResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkBindClientTelResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkBindClientTelResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkBindClientTelResponseBody) GetData() *ClinkBindClientTelResponseBodyData {
	return s.Data
}

func (s *ClinkBindClientTelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkBindClientTelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkBindClientTelResponseBody) SetAccessDeniedDetail(v string) *ClinkBindClientTelResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkBindClientTelResponseBody) SetCode(v string) *ClinkBindClientTelResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkBindClientTelResponseBody) SetData(v *ClinkBindClientTelResponseBodyData) *ClinkBindClientTelResponseBody {
	s.Data = v
	return s
}

func (s *ClinkBindClientTelResponseBody) SetMessage(v string) *ClinkBindClientTelResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkBindClientTelResponseBody) SetRequestId(v string) *ClinkBindClientTelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkBindClientTelResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkBindClientTelResponseBodyData struct {
	// 请求 id
	//
	// example:
	//
	// null
	ClinkRequestId *string `json:"ClinkRequestId,omitempty" xml:"ClinkRequestId,omitempty"`
}

func (s ClinkBindClientTelResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkBindClientTelResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkBindClientTelResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkBindClientTelResponseBodyData) SetClinkRequestId(v string) *ClinkBindClientTelResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkBindClientTelResponseBodyData) Validate() error {
	return dara.Validate(s)
}
