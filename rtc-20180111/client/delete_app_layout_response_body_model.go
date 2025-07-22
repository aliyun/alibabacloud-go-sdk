// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppLayoutResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAppLayoutResponseBody
	GetRequestId() *string
}

type DeleteAppLayoutResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 2DCE8D7E-BE3B-54AB-8DAC-32F34BED0763
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppLayoutResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppLayoutResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppLayoutResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppLayoutResponseBody) SetRequestId(v string) *DeleteAppLayoutResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppLayoutResponseBody) Validate() error {
	return dara.Validate(s)
}
