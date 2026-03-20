// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProjectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteProjectResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteProjectResponseBody
	GetSuccess() *bool
}

type DeleteProjectResponseBody struct {
	// example:
	//
	// 20260128101840d2a3770b05d299b3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteProjectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteProjectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteProjectResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteProjectResponseBody) SetRequestId(v string) *DeleteProjectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteProjectResponseBody) SetSuccess(v bool) *DeleteProjectResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteProjectResponseBody) Validate() error {
	return dara.Validate(s)
}
