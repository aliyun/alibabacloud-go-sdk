// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAdvancedSearchFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAdvancedSearchFileResponseBody
	GetRequestId() *string
}

type CreateAdvancedSearchFileResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D231366D-E2AD-559E-9C29-58FF7F4B1F5D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAdvancedSearchFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAdvancedSearchFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAdvancedSearchFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAdvancedSearchFileResponseBody) SetRequestId(v string) *CreateAdvancedSearchFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAdvancedSearchFileResponseBody) Validate() error {
	return dara.Validate(s)
}
