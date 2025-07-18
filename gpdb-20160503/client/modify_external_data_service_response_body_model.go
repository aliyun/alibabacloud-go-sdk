// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyExternalDataServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyExternalDataServiceResponseBody
	GetRequestId() *string
}

type ModifyExternalDataServiceResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyExternalDataServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyExternalDataServiceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyExternalDataServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyExternalDataServiceResponseBody) SetRequestId(v string) *ModifyExternalDataServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyExternalDataServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
