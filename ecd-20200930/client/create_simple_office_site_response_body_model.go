// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSimpleOfficeSiteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOfficeSiteId(v string) *CreateSimpleOfficeSiteResponseBody
	GetOfficeSiteId() *string
	SetRequestId(v string) *CreateSimpleOfficeSiteResponseBody
	GetRequestId() *string
}

type CreateSimpleOfficeSiteResponseBody struct {
	// The office network ID.
	//
	// example:
	//
	// cn-hangzhou+os-c5cy7q578s8jc****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSimpleOfficeSiteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSimpleOfficeSiteResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSimpleOfficeSiteResponseBody) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *CreateSimpleOfficeSiteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSimpleOfficeSiteResponseBody) SetOfficeSiteId(v string) *CreateSimpleOfficeSiteResponseBody {
	s.OfficeSiteId = &v
	return s
}

func (s *CreateSimpleOfficeSiteResponseBody) SetRequestId(v string) *CreateSimpleOfficeSiteResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSimpleOfficeSiteResponseBody) Validate() error {
	return dara.Validate(s)
}
