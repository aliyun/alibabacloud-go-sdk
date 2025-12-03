// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEaiEcsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClientInstanceId(v string) *CreateEaiEcsResponseBody
	GetClientInstanceId() *string
	SetElasticAcceleratedInstanceId(v string) *CreateEaiEcsResponseBody
	GetElasticAcceleratedInstanceId() *string
	SetRequestId(v string) *CreateEaiEcsResponseBody
	GetRequestId() *string
}

type CreateEaiEcsResponseBody struct {
	// example:
	//
	// i-bp1hjrvleawl4ogb****
	ClientInstanceId *string `json:"ClientInstanceId,omitempty" xml:"ClientInstanceId,omitempty"`
	// example:
	//
	// eais-sz8t15a7gt7z7j7i****
	ElasticAcceleratedInstanceId *string `json:"ElasticAcceleratedInstanceId,omitempty" xml:"ElasticAcceleratedInstanceId,omitempty"`
	// example:
	//
	// F5FEB9AA-C108-577C-AB3D-D13524AF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEaiEcsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEaiEcsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEaiEcsResponseBody) GetClientInstanceId() *string {
	return s.ClientInstanceId
}

func (s *CreateEaiEcsResponseBody) GetElasticAcceleratedInstanceId() *string {
	return s.ElasticAcceleratedInstanceId
}

func (s *CreateEaiEcsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEaiEcsResponseBody) SetClientInstanceId(v string) *CreateEaiEcsResponseBody {
	s.ClientInstanceId = &v
	return s
}

func (s *CreateEaiEcsResponseBody) SetElasticAcceleratedInstanceId(v string) *CreateEaiEcsResponseBody {
	s.ElasticAcceleratedInstanceId = &v
	return s
}

func (s *CreateEaiEcsResponseBody) SetRequestId(v string) *CreateEaiEcsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEaiEcsResponseBody) Validate() error {
	return dara.Validate(s)
}
