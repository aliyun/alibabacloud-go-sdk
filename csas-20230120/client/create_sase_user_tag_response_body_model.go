// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSaseUserTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSaseUserTagResponseBody
	GetRequestId() *string
	SetTagId(v string) *CreateSaseUserTagResponseBody
	GetTagId() *string
}

type CreateSaseUserTagResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FEF1144C-95D1-5F7C-81EF-9DB70EA49FCE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The user tag ID.
	//
	// example:
	//
	// su-tag-1ae52f66039fa0d4****
	TagId *string `json:"TagId,omitempty" xml:"TagId,omitempty"`
}

func (s CreateSaseUserTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSaseUserTagResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSaseUserTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSaseUserTagResponseBody) GetTagId() *string {
	return s.TagId
}

func (s *CreateSaseUserTagResponseBody) SetRequestId(v string) *CreateSaseUserTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSaseUserTagResponseBody) SetTagId(v string) *CreateSaseUserTagResponseBody {
	s.TagId = &v
	return s
}

func (s *CreateSaseUserTagResponseBody) Validate() error {
	return dara.Validate(s)
}
