// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantInstanceToVbrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GrantInstanceToVbrResponseBody
	GetRequestId() *string
}

type GrantInstanceToVbrResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F99F13AE-D733-5856-AB97-80CC88B1D5A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GrantInstanceToVbrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantInstanceToVbrResponseBody) GoString() string {
	return s.String()
}

func (s *GrantInstanceToVbrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantInstanceToVbrResponseBody) SetRequestId(v string) *GrantInstanceToVbrResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantInstanceToVbrResponseBody) Validate() error {
	return dara.Validate(s)
}
