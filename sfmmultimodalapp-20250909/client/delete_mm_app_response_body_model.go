// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMmAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMmAppResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMmAppResponseBody
	GetSuccess() *bool
}

type DeleteMmAppResponseBody struct {
	// example:
	//
	// xxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMmAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMmAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMmAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMmAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMmAppResponseBody) SetRequestId(v string) *DeleteMmAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMmAppResponseBody) SetSuccess(v bool) *DeleteMmAppResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMmAppResponseBody) Validate() error {
	return dara.Validate(s)
}
