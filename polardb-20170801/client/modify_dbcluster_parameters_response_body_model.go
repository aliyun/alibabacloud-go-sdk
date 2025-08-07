// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterParametersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBClusterParametersResponseBody
	GetRequestId() *string
}

type ModifyDBClusterParametersResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C5D526E5-91B5-48B8-B980-FE07FF******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterParametersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterParametersResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterParametersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterParametersResponseBody) SetRequestId(v string) *ModifyDBClusterParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterParametersResponseBody) Validate() error {
	return dara.Validate(s)
}
