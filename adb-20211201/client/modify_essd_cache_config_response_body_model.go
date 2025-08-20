// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyEssdCacheConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyEssdCacheConfigResponseBody
	GetRequestId() *string
}

type ModifyEssdCacheConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 115F9CCA-EF2E-5F91-AB60-4961D52FEAB4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyEssdCacheConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyEssdCacheConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyEssdCacheConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyEssdCacheConfigResponseBody) SetRequestId(v string) *ModifyEssdCacheConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyEssdCacheConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
