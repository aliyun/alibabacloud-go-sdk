// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEaiJupyterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetElasticAcceleratedInstanceId(v string) *CreateEaiJupyterResponseBody
	GetElasticAcceleratedInstanceId() *string
	SetRequestId(v string) *CreateEaiJupyterResponseBody
	GetRequestId() *string
}

type CreateEaiJupyterResponseBody struct {
	// example:
	//
	// eais-hz8t15a7gt7z7j7i****
	ElasticAcceleratedInstanceId *string `json:"ElasticAcceleratedInstanceId,omitempty" xml:"ElasticAcceleratedInstanceId,omitempty"`
	// example:
	//
	// A655AB0E-31BB-45AD-9255-FCE93F6*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEaiJupyterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEaiJupyterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEaiJupyterResponseBody) GetElasticAcceleratedInstanceId() *string {
	return s.ElasticAcceleratedInstanceId
}

func (s *CreateEaiJupyterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEaiJupyterResponseBody) SetElasticAcceleratedInstanceId(v string) *CreateEaiJupyterResponseBody {
	s.ElasticAcceleratedInstanceId = &v
	return s
}

func (s *CreateEaiJupyterResponseBody) SetRequestId(v string) *CreateEaiJupyterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEaiJupyterResponseBody) Validate() error {
	return dara.Validate(s)
}
