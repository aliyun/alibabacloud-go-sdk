// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushDocumentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v []interface{}) *PushDocumentsRequest
	GetBody() []interface{}
	SetPkField(v string) *PushDocumentsRequest
	GetPkField() *string
}

type PushDocumentsRequest struct {
	// The request body.
	Body []interface{} `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
	// The primary key field.
	//
	// example:
	//
	// id
	PkField *string `json:"pkField,omitempty" xml:"pkField,omitempty"`
}

func (s PushDocumentsRequest) String() string {
	return dara.Prettify(s)
}

func (s PushDocumentsRequest) GoString() string {
	return s.String()
}

func (s *PushDocumentsRequest) GetBody() []interface{} {
	return s.Body
}

func (s *PushDocumentsRequest) GetPkField() *string {
	return s.PkField
}

func (s *PushDocumentsRequest) SetBody(v []interface{}) *PushDocumentsRequest {
	s.Body = v
	return s
}

func (s *PushDocumentsRequest) SetPkField(v string) *PushDocumentsRequest {
	s.PkField = &v
	return s
}

func (s *PushDocumentsRequest) Validate() error {
	return dara.Validate(s)
}
