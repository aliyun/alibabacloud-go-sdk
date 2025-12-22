// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFunctionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFunctions(v []*string) *ListFunctionsResponseBody
	GetFunctions() []*string
	SetNextPageToken(v string) *ListFunctionsResponseBody
	GetNextPageToken() *string
}

type ListFunctionsResponseBody struct {
	Functions []*string `json:"functions,omitempty" xml:"functions,omitempty" type:"Repeated"`
	// example:
	//
	// E8ABEB1C3DB893D16576269017992F57
	NextPageToken *string `json:"nextPageToken,omitempty" xml:"nextPageToken,omitempty"`
}

func (s ListFunctionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFunctionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFunctionsResponseBody) GetFunctions() []*string {
	return s.Functions
}

func (s *ListFunctionsResponseBody) GetNextPageToken() *string {
	return s.NextPageToken
}

func (s *ListFunctionsResponseBody) SetFunctions(v []*string) *ListFunctionsResponseBody {
	s.Functions = v
	return s
}

func (s *ListFunctionsResponseBody) SetNextPageToken(v string) *ListFunctionsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListFunctionsResponseBody) Validate() error {
	return dara.Validate(s)
}
