// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcHoneyPotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVpcHoneyPotResponseBody
	GetRequestId() *string
}

type DeleteVpcHoneyPotResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 4347E985-6E64-467B-96EC-30D4EA9E32FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVpcHoneyPotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcHoneyPotResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpcHoneyPotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVpcHoneyPotResponseBody) SetRequestId(v string) *DeleteVpcHoneyPotResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVpcHoneyPotResponseBody) Validate() error {
	return dara.Validate(s)
}
