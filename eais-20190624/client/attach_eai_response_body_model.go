// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachEaiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClientInstanceId(v string) *AttachEaiResponseBody
	GetClientInstanceId() *string
	SetElasticAcceleratedInstanceId(v string) *AttachEaiResponseBody
	GetElasticAcceleratedInstanceId() *string
	SetRequestId(v string) *AttachEaiResponseBody
	GetRequestId() *string
}

type AttachEaiResponseBody struct {
	// example:
	//
	// i-wz93g6pyat2g7t7o****
	ClientInstanceId *string `json:"ClientInstanceId,omitempty" xml:"ClientInstanceId,omitempty"`
	// example:
	//
	// eais-sz8t15a7gt7z7j7i****
	ElasticAcceleratedInstanceId *string `json:"ElasticAcceleratedInstanceId,omitempty" xml:"ElasticAcceleratedInstanceId,omitempty"`
	// example:
	//
	// C3BCB7DA-BEB6-4982-A765-6EA61EC84474
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachEaiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachEaiResponseBody) GoString() string {
	return s.String()
}

func (s *AttachEaiResponseBody) GetClientInstanceId() *string {
	return s.ClientInstanceId
}

func (s *AttachEaiResponseBody) GetElasticAcceleratedInstanceId() *string {
	return s.ElasticAcceleratedInstanceId
}

func (s *AttachEaiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachEaiResponseBody) SetClientInstanceId(v string) *AttachEaiResponseBody {
	s.ClientInstanceId = &v
	return s
}

func (s *AttachEaiResponseBody) SetElasticAcceleratedInstanceId(v string) *AttachEaiResponseBody {
	s.ElasticAcceleratedInstanceId = &v
	return s
}

func (s *AttachEaiResponseBody) SetRequestId(v string) *AttachEaiResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachEaiResponseBody) Validate() error {
	return dara.Validate(s)
}
