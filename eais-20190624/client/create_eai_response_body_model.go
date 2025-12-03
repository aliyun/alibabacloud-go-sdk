// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEaiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetElasticAcceleratedInstanceId(v string) *CreateEaiResponseBody
	GetElasticAcceleratedInstanceId() *string
	SetRequestId(v string) *CreateEaiResponseBody
	GetRequestId() *string
}

type CreateEaiResponseBody struct {
	// example:
	//
	// eais-sz8t15a7gt7z7j7i****
	ElasticAcceleratedInstanceId *string `json:"ElasticAcceleratedInstanceId,omitempty" xml:"ElasticAcceleratedInstanceId,omitempty"`
	// example:
	//
	// A655AB0E-31BB-45AD-9255-FCE93F6*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEaiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEaiResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEaiResponseBody) GetElasticAcceleratedInstanceId() *string {
	return s.ElasticAcceleratedInstanceId
}

func (s *CreateEaiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEaiResponseBody) SetElasticAcceleratedInstanceId(v string) *CreateEaiResponseBody {
	s.ElasticAcceleratedInstanceId = &v
	return s
}

func (s *CreateEaiResponseBody) SetRequestId(v string) *CreateEaiResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEaiResponseBody) Validate() error {
	return dara.Validate(s)
}
