// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateOfficeSiteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ActivateOfficeSiteResponseBody
	GetRequestId() *string
}

type ActivateOfficeSiteResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 269BDB16-2CD8-4865-84BD-11C40BC21DB0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ActivateOfficeSiteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ActivateOfficeSiteResponseBody) GoString() string {
	return s.String()
}

func (s *ActivateOfficeSiteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ActivateOfficeSiteResponseBody) SetRequestId(v string) *ActivateOfficeSiteResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActivateOfficeSiteResponseBody) Validate() error {
	return dara.Validate(s)
}
