// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlgorithmDefResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAlgorithmDefResponseBody
	GetRequestId() *string
	SetSpec(v map[string]interface{}) *GetAlgorithmDefResponseBody
	GetSpec() map[string]interface{}
}

type GetAlgorithmDefResponseBody struct {
	// example:
	//
	// B4F16666-FD54-5D9D-A362-53A4C66692DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// {}
	Spec map[string]interface{} `json:"Spec,omitempty" xml:"Spec,omitempty"`
}

func (s GetAlgorithmDefResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAlgorithmDefResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlgorithmDefResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAlgorithmDefResponseBody) GetSpec() map[string]interface{} {
	return s.Spec
}

func (s *GetAlgorithmDefResponseBody) SetRequestId(v string) *GetAlgorithmDefResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlgorithmDefResponseBody) SetSpec(v map[string]interface{}) *GetAlgorithmDefResponseBody {
	s.Spec = v
	return s
}

func (s *GetAlgorithmDefResponseBody) Validate() error {
	return dara.Validate(s)
}
