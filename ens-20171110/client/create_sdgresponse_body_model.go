// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSDGResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSDGResponseBody
	GetRequestId() *string
	SetSDGId(v string) *CreateSDGResponseBody
	GetSDGId() *string
}

type CreateSDGResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the generated SDG.
	//
	// example:
	//
	// sdg-xxxx
	SDGId *string `json:"SDGId,omitempty" xml:"SDGId,omitempty"`
}

func (s CreateSDGResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSDGResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSDGResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSDGResponseBody) GetSDGId() *string {
	return s.SDGId
}

func (s *CreateSDGResponseBody) SetRequestId(v string) *CreateSDGResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSDGResponseBody) SetSDGId(v string) *CreateSDGResponseBody {
	s.SDGId = &v
	return s
}

func (s *CreateSDGResponseBody) Validate() error {
	return dara.Validate(s)
}
