// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultRegistrantProfileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDefaultRegistrantProfileResponseBody
	GetRequestId() *string
}

type SetDefaultRegistrantProfileResponseBody struct {
	// example:
	//
	// 4D73432C-7600-4779-ACBB-C3B5CA145D32
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDefaultRegistrantProfileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultRegistrantProfileResponseBody) GoString() string {
	return s.String()
}

func (s *SetDefaultRegistrantProfileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDefaultRegistrantProfileResponseBody) SetRequestId(v string) *SetDefaultRegistrantProfileResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDefaultRegistrantProfileResponseBody) Validate() error {
	return dara.Validate(s)
}
