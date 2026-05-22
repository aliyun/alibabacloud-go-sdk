// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSlrRoleForRealtimeLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSlrRoleForRealtimeLogResponseBody
	GetRequestId() *string
}

type CreateSlrRoleForRealtimeLogResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 156A6B-677B1A-4297B7-9187B7-2B44792
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateSlrRoleForRealtimeLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSlrRoleForRealtimeLogResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSlrRoleForRealtimeLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSlrRoleForRealtimeLogResponseBody) SetRequestId(v string) *CreateSlrRoleForRealtimeLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSlrRoleForRealtimeLogResponseBody) Validate() error {
	return dara.Validate(s)
}
