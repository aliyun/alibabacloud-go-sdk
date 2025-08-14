// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventFieldResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEventFieldResponseBody
	GetRequestId() *string
	SetResuleObject(v bool) *DeleteEventFieldResponseBody
	GetResuleObject() *bool
}

type DeleteEventFieldResponseBody struct {
	// ID of the request
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Result object
	//
	// example:
	//
	// true
	ResuleObject *bool `json:"resuleObject,omitempty" xml:"resuleObject,omitempty"`
}

func (s DeleteEventFieldResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventFieldResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventFieldResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEventFieldResponseBody) GetResuleObject() *bool {
	return s.ResuleObject
}

func (s *DeleteEventFieldResponseBody) SetRequestId(v string) *DeleteEventFieldResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEventFieldResponseBody) SetResuleObject(v bool) *DeleteEventFieldResponseBody {
	s.ResuleObject = &v
	return s
}

func (s *DeleteEventFieldResponseBody) Validate() error {
	return dara.Validate(s)
}
