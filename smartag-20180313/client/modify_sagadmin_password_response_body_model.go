// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySAGAdminPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySAGAdminPasswordResponseBody
	GetRequestId() *string
}

type ModifySAGAdminPasswordResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// DB0A026C-A8E5-40AB-977E-3A87DD78F694
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySAGAdminPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySAGAdminPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySAGAdminPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySAGAdminPasswordResponseBody) SetRequestId(v string) *ModifySAGAdminPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySAGAdminPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
