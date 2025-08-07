// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBDescriptionResponseBody
	GetRequestId() *string
}

type ModifyDBDescriptionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5F295AAB-B599-4F0E-96D0-DBFAB3******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBDescriptionResponseBody) SetRequestId(v string) *ModifyDBDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
