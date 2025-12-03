// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachEaisEiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClientInstanceId(v string) *AttachEaisEiResponseBody
	GetClientInstanceId() *string
	SetEiInstanceId(v string) *AttachEaisEiResponseBody
	GetEiInstanceId() *string
	SetRequestId(v string) *AttachEaisEiResponseBody
	GetRequestId() *string
}

type AttachEaisEiResponseBody struct {
	// example:
	//
	// i-bp14ws9hbt1oe0u9****
	ClientInstanceId *string `json:"ClientInstanceId,omitempty" xml:"ClientInstanceId,omitempty"`
	// example:
	//
	// eais-hzu00xufs1c8j5nn****
	EiInstanceId *string `json:"EiInstanceId,omitempty" xml:"EiInstanceId,omitempty"`
	// example:
	//
	// C3BCB7DA-BEB6-4982-A765-6EA61EC8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachEaisEiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachEaisEiResponseBody) GoString() string {
	return s.String()
}

func (s *AttachEaisEiResponseBody) GetClientInstanceId() *string {
	return s.ClientInstanceId
}

func (s *AttachEaisEiResponseBody) GetEiInstanceId() *string {
	return s.EiInstanceId
}

func (s *AttachEaisEiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachEaisEiResponseBody) SetClientInstanceId(v string) *AttachEaisEiResponseBody {
	s.ClientInstanceId = &v
	return s
}

func (s *AttachEaisEiResponseBody) SetEiInstanceId(v string) *AttachEaisEiResponseBody {
	s.EiInstanceId = &v
	return s
}

func (s *AttachEaisEiResponseBody) SetRequestId(v string) *AttachEaisEiResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachEaisEiResponseBody) Validate() error {
	return dara.Validate(s)
}
