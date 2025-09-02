// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPublicModelEngineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *QueryPublicModelEngineResponseBody
	GetRequestId() *string
	SetReturnValue(v []map[string]interface{}) *QueryPublicModelEngineResponseBody
	GetReturnValue() []map[string]interface{}
}

type QueryPublicModelEngineResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4CFB2DED-7D9B-4C42-B4AA-CFF4991DFFF4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned information about objects.
	ReturnValue []map[string]interface{} `json:"ReturnValue,omitempty" xml:"ReturnValue,omitempty" type:"Repeated"`
}

func (s QueryPublicModelEngineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPublicModelEngineResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPublicModelEngineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPublicModelEngineResponseBody) GetReturnValue() []map[string]interface{} {
	return s.ReturnValue
}

func (s *QueryPublicModelEngineResponseBody) SetRequestId(v string) *QueryPublicModelEngineResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPublicModelEngineResponseBody) SetReturnValue(v []map[string]interface{}) *QueryPublicModelEngineResponseBody {
	s.ReturnValue = v
	return s
}

func (s *QueryPublicModelEngineResponseBody) Validate() error {
	return dara.Validate(s)
}
