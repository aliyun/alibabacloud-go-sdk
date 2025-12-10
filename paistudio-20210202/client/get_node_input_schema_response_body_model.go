// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeInputSchemaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetColNames(v []*string) *GetNodeInputSchemaResponseBody
	GetColNames() []*string
	SetColTypes(v []*string) *GetNodeInputSchemaResponseBody
	GetColTypes() []*string
	SetRequestId(v string) *GetNodeInputSchemaResponseBody
	GetRequestId() *string
}

type GetNodeInputSchemaResponseBody struct {
	ColNames []*string `json:"ColNames,omitempty" xml:"ColNames,omitempty" type:"Repeated"`
	ColTypes []*string `json:"ColTypes,omitempty" xml:"ColTypes,omitempty" type:"Repeated"`
	// example:
	//
	// CEB07647-8A5D-56F1-8B99-361BCF51402F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetNodeInputSchemaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetNodeInputSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *GetNodeInputSchemaResponseBody) GetColNames() []*string {
	return s.ColNames
}

func (s *GetNodeInputSchemaResponseBody) GetColTypes() []*string {
	return s.ColTypes
}

func (s *GetNodeInputSchemaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetNodeInputSchemaResponseBody) SetColNames(v []*string) *GetNodeInputSchemaResponseBody {
	s.ColNames = v
	return s
}

func (s *GetNodeInputSchemaResponseBody) SetColTypes(v []*string) *GetNodeInputSchemaResponseBody {
	s.ColTypes = v
	return s
}

func (s *GetNodeInputSchemaResponseBody) SetRequestId(v string) *GetNodeInputSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetNodeInputSchemaResponseBody) Validate() error {
	return dara.Validate(s)
}
