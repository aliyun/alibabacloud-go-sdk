// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBNodeConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBNodeConfigResponseBody
	GetRequestId() *string
}

type ModifyDBNodeConfigResponseBody struct {
	// example:
	//
	// 925B84D9-CA72-432C-95CF-738C22******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBNodeConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBNodeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBNodeConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBNodeConfigResponseBody) SetRequestId(v string) *ModifyDBNodeConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBNodeConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
