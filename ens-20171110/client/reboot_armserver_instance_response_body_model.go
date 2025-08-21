// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootARMServerInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RebootARMServerInstanceResponseBody
	GetRequestId() *string
}

type RebootARMServerInstanceResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebootARMServerInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RebootARMServerInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RebootARMServerInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebootARMServerInstanceResponseBody) SetRequestId(v string) *RebootARMServerInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebootARMServerInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
