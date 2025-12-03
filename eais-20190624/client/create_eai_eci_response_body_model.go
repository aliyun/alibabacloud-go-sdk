// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEaiEciResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClientInstanceId(v string) *CreateEaiEciResponseBody
	GetClientInstanceId() *string
	SetElasticAcceleratedInstanceId(v string) *CreateEaiEciResponseBody
	GetElasticAcceleratedInstanceId() *string
	SetRequestId(v string) *CreateEaiEciResponseBody
	GetRequestId() *string
}

type CreateEaiEciResponseBody struct {
	// example:
	//
	// eci-2zeh03ygxlrzmfi6****
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

func (s CreateEaiEciResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEaiEciResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEaiEciResponseBody) GetClientInstanceId() *string {
	return s.ClientInstanceId
}

func (s *CreateEaiEciResponseBody) GetElasticAcceleratedInstanceId() *string {
	return s.ElasticAcceleratedInstanceId
}

func (s *CreateEaiEciResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEaiEciResponseBody) SetClientInstanceId(v string) *CreateEaiEciResponseBody {
	s.ClientInstanceId = &v
	return s
}

func (s *CreateEaiEciResponseBody) SetElasticAcceleratedInstanceId(v string) *CreateEaiEciResponseBody {
	s.ElasticAcceleratedInstanceId = &v
	return s
}

func (s *CreateEaiEciResponseBody) SetRequestId(v string) *CreateEaiEciResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEaiEciResponseBody) Validate() error {
	return dara.Validate(s)
}
