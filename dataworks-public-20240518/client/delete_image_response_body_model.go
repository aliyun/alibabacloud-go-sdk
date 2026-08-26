// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteImageResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteImageResponseBody
	GetSuccess() *bool
}

type DeleteImageResponseBody struct {
	// The result of the API request.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID, which is used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteImageResponseBody) SetData(v bool) *DeleteImageResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteImageResponseBody) SetRequestId(v string) *DeleteImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteImageResponseBody) SetSuccess(v bool) *DeleteImageResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteImageResponseBody) Validate() error {
	return dara.Validate(s)
}
