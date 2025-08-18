// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCnameFlatteningResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFlattenMode(v string) *GetCnameFlatteningResponseBody
	GetFlattenMode() *string
	SetRequestId(v string) *GetCnameFlatteningResponseBody
	GetRequestId() *string
}

type GetCnameFlatteningResponseBody struct {
	// The CNAME flattening mode. Valid values:
	//
	// 	- flatten_all: flattens all CNAMEs.
	//
	// 	- flatten_all (default): flattens only the root domain.
	//
	// example:
	//
	// flatten_all
	FlattenMode *string `json:"FlattenMode,omitempty" xml:"FlattenMode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F61CDR30-E83C-4FDA-BF73-9A94CDD44229
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCnameFlatteningResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCnameFlatteningResponseBody) GoString() string {
	return s.String()
}

func (s *GetCnameFlatteningResponseBody) GetFlattenMode() *string {
	return s.FlattenMode
}

func (s *GetCnameFlatteningResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCnameFlatteningResponseBody) SetFlattenMode(v string) *GetCnameFlatteningResponseBody {
	s.FlattenMode = &v
	return s
}

func (s *GetCnameFlatteningResponseBody) SetRequestId(v string) *GetCnameFlatteningResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCnameFlatteningResponseBody) Validate() error {
	return dara.Validate(s)
}
