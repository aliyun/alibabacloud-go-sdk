// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICenterStateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAICenterStateResponseBody
	GetRequestId() *string
	SetState(v string) *GetAICenterStateResponseBody
	GetState() *string
}

type GetAICenterStateResponseBody struct {
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// Running
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s GetAICenterStateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAICenterStateResponseBody) GoString() string {
	return s.String()
}

func (s *GetAICenterStateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAICenterStateResponseBody) GetState() *string {
	return s.State
}

func (s *GetAICenterStateResponseBody) SetRequestId(v string) *GetAICenterStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAICenterStateResponseBody) SetState(v string) *GetAICenterStateResponseBody {
	s.State = &v
	return s
}

func (s *GetAICenterStateResponseBody) Validate() error {
	return dara.Validate(s)
}
