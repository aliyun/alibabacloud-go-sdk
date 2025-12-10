// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlgorithmDefsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAlgorithmDefsResponseBody
	GetRequestId() *string
	SetSpecs(v []map[string]interface{}) *GetAlgorithmDefsResponseBody
	GetSpecs() []map[string]interface{}
}

type GetAlgorithmDefsResponseBody struct {
	// example:
	//
	// F082BD0D-21E1-5F9B-81A0-AB07485B03CD
	RequestId *string                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Specs     []map[string]interface{} `json:"Specs,omitempty" xml:"Specs,omitempty" type:"Repeated"`
}

func (s GetAlgorithmDefsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAlgorithmDefsResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlgorithmDefsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAlgorithmDefsResponseBody) GetSpecs() []map[string]interface{} {
	return s.Specs
}

func (s *GetAlgorithmDefsResponseBody) SetRequestId(v string) *GetAlgorithmDefsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlgorithmDefsResponseBody) SetSpecs(v []map[string]interface{}) *GetAlgorithmDefsResponseBody {
	s.Specs = v
	return s
}

func (s *GetAlgorithmDefsResponseBody) Validate() error {
	return dara.Validate(s)
}
