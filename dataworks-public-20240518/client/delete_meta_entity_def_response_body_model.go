// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMetaEntityDefResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMetaEntityDefResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteMetaEntityDefResponseBody
	GetSuccess() *bool
}

type DeleteMetaEntityDefResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 3C1E755D-B606-57A4-9B9C-7B214E81354C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMetaEntityDefResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMetaEntityDefResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMetaEntityDefResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMetaEntityDefResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteMetaEntityDefResponseBody) SetRequestId(v string) *DeleteMetaEntityDefResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMetaEntityDefResponseBody) SetSuccess(v bool) *DeleteMetaEntityDefResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMetaEntityDefResponseBody) Validate() error {
	return dara.Validate(s)
}
