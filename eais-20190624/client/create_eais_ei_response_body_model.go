// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEaisEiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEiInstanceId(v string) *CreateEaisEiResponseBody
	GetEiInstanceId() *string
	SetRequestId(v string) *CreateEaisEiResponseBody
	GetRequestId() *string
}

type CreateEaisEiResponseBody struct {
	// example:
	//
	// eais-hzu00xufs1c8j5nn****
	EiInstanceId *string `json:"EiInstanceId,omitempty" xml:"EiInstanceId,omitempty"`
	// example:
	//
	// F5FEB9AA-C108-577C-AB3D-D13524AF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEaisEiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEaisEiResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEaisEiResponseBody) GetEiInstanceId() *string {
	return s.EiInstanceId
}

func (s *CreateEaisEiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEaisEiResponseBody) SetEiInstanceId(v string) *CreateEaisEiResponseBody {
	s.EiInstanceId = &v
	return s
}

func (s *CreateEaisEiResponseBody) SetRequestId(v string) *CreateEaisEiResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEaisEiResponseBody) Validate() error {
	return dara.Validate(s)
}
