// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteParentPlatformResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteParentPlatformResponseBody
	GetRequestId() *string
}

type DeleteParentPlatformResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteParentPlatformResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteParentPlatformResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteParentPlatformResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteParentPlatformResponseBody) SetRequestId(v string) *DeleteParentPlatformResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteParentPlatformResponseBody) Validate() error {
	return dara.Validate(s)
}
