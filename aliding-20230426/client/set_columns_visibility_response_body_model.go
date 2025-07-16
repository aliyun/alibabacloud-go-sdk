// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetColumnsVisibilityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *SetColumnsVisibilityResponseBody
	GetId() *string
	SetRequestId(v string) *SetColumnsVisibilityResponseBody
	GetRequestId() *string
}

type SetColumnsVisibilityResponseBody struct {
	// example:
	//
	// stxxxx
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// requestId
	//
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SetColumnsVisibilityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetColumnsVisibilityResponseBody) GoString() string {
	return s.String()
}

func (s *SetColumnsVisibilityResponseBody) GetId() *string {
	return s.Id
}

func (s *SetColumnsVisibilityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetColumnsVisibilityResponseBody) SetId(v string) *SetColumnsVisibilityResponseBody {
	s.Id = &v
	return s
}

func (s *SetColumnsVisibilityResponseBody) SetRequestId(v string) *SetColumnsVisibilityResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetColumnsVisibilityResponseBody) Validate() error {
	return dara.Validate(s)
}
