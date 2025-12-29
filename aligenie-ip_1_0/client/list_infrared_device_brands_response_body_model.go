// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInfraredDeviceBrandsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ListInfraredDeviceBrandsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListInfraredDeviceBrandsResponseBody
	GetRequestId() *string
	SetResult(v map[string][]*string) *ListInfraredDeviceBrandsResponseBody
	GetResult() map[string][]*string
	SetStatusCode(v int32) *ListInfraredDeviceBrandsResponseBody
	GetStatusCode() *int32
}

type ListInfraredDeviceBrandsResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 860194F7-9593-50EA-8E53-BCEC0D325A00
	RequestId *string              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    map[string][]*string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ListInfraredDeviceBrandsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInfraredDeviceBrandsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInfraredDeviceBrandsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListInfraredDeviceBrandsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInfraredDeviceBrandsResponseBody) GetResult() map[string][]*string {
	return s.Result
}

func (s *ListInfraredDeviceBrandsResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInfraredDeviceBrandsResponseBody) SetMessage(v string) *ListInfraredDeviceBrandsResponseBody {
	s.Message = &v
	return s
}

func (s *ListInfraredDeviceBrandsResponseBody) SetRequestId(v string) *ListInfraredDeviceBrandsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInfraredDeviceBrandsResponseBody) SetResult(v map[string][]*string) *ListInfraredDeviceBrandsResponseBody {
	s.Result = v
	return s
}

func (s *ListInfraredDeviceBrandsResponseBody) SetStatusCode(v int32) *ListInfraredDeviceBrandsResponseBody {
	s.StatusCode = &v
	return s
}

func (s *ListInfraredDeviceBrandsResponseBody) Validate() error {
	return dara.Validate(s)
}
