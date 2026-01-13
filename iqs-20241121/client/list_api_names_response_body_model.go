// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiNamesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiNames(v []*string) *ListApiNamesResponseBody
	GetApiNames() []*string
	SetRequestId(v string) *ListApiNamesResponseBody
	GetRequestId() *string
}

type ListApiNamesResponseBody struct {
	ApiNames []*string `json:"apiNames,omitempty" xml:"apiNames,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// ECB2144C-E277-5434-80E6-12D26678D364
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListApiNamesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApiNamesResponseBody) GoString() string {
	return s.String()
}

func (s *ListApiNamesResponseBody) GetApiNames() []*string {
	return s.ApiNames
}

func (s *ListApiNamesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApiNamesResponseBody) SetApiNames(v []*string) *ListApiNamesResponseBody {
	s.ApiNames = v
	return s
}

func (s *ListApiNamesResponseBody) SetRequestId(v string) *ListApiNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApiNamesResponseBody) Validate() error {
	return dara.Validate(s)
}
