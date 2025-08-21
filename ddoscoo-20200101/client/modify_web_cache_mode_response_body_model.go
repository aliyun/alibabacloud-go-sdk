// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebCacheModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyWebCacheModeResponseBody
	GetRequestId() *string
}

type ModifyWebCacheModeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebCacheModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebCacheModeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebCacheModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyWebCacheModeResponseBody) SetRequestId(v string) *ModifyWebCacheModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWebCacheModeResponseBody) Validate() error {
	return dara.Validate(s)
}
