// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDiagnoseSmartAccessGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DiagnoseSmartAccessGatewayResponseBody
	GetRequestId() *string
}

type DiagnoseSmartAccessGatewayResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 193AE392-76C2-4D3E-9420-889A51B43CC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DiagnoseSmartAccessGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DiagnoseSmartAccessGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *DiagnoseSmartAccessGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DiagnoseSmartAccessGatewayResponseBody) SetRequestId(v string) *DiagnoseSmartAccessGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *DiagnoseSmartAccessGatewayResponseBody) Validate() error {
	return dara.Validate(s)
}
