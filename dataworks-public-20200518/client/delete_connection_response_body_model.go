// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteConnectionResponseBody
	GetData() *bool
	SetHttpStatusCode(v string) *DeleteConnectionResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *DeleteConnectionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteConnectionResponseBody
	GetSuccess() *bool
}

type DeleteConnectionResponseBody struct {
	// Indicates whether the data source was removed.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0bc1411515937635****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConnectionResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteConnectionResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DeleteConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteConnectionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteConnectionResponseBody) SetData(v bool) *DeleteConnectionResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteConnectionResponseBody) SetHttpStatusCode(v string) *DeleteConnectionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteConnectionResponseBody) SetRequestId(v string) *DeleteConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConnectionResponseBody) SetSuccess(v bool) *DeleteConnectionResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
