// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConnectorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteConnectorResponseBody
	GetRequestId() *string
}

type DeleteConnectorResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// EA4D25BE-BBAB-553E-B18C-32976CFDE86B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteConnectorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConnectorResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConnectorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteConnectorResponseBody) SetRequestId(v string) *DeleteConnectorResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConnectorResponseBody) Validate() error {
	return dara.Validate(s)
}
