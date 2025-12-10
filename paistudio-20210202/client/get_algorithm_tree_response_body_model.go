// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlgorithmTreeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetAlgorithmTreeResponseBody
	GetRequestId() *string
	SetTimestamp(v string) *GetAlgorithmTreeResponseBody
	GetTimestamp() *string
	SetTree(v []map[string]interface{}) *GetAlgorithmTreeResponseBody
	GetTree() []map[string]interface{}
}

type GetAlgorithmTreeResponseBody struct {
	// example:
	//
	// 46B59732-033F-5C96-9B15-1E05E7705548
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20210317101023
	Timestamp *string                  `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	Tree      []map[string]interface{} `json:"Tree,omitempty" xml:"Tree,omitempty" type:"Repeated"`
}

func (s GetAlgorithmTreeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAlgorithmTreeResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlgorithmTreeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAlgorithmTreeResponseBody) GetTimestamp() *string {
	return s.Timestamp
}

func (s *GetAlgorithmTreeResponseBody) GetTree() []map[string]interface{} {
	return s.Tree
}

func (s *GetAlgorithmTreeResponseBody) SetRequestId(v string) *GetAlgorithmTreeResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlgorithmTreeResponseBody) SetTimestamp(v string) *GetAlgorithmTreeResponseBody {
	s.Timestamp = &v
	return s
}

func (s *GetAlgorithmTreeResponseBody) SetTree(v []map[string]interface{}) *GetAlgorithmTreeResponseBody {
	s.Tree = v
	return s
}

func (s *GetAlgorithmTreeResponseBody) Validate() error {
	return dara.Validate(s)
}
