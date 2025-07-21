// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppOtaVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAppOtaVersionsResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteAppOtaVersionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAppOtaVersionsResponseBody
	GetRequestId() *string
}

type DeleteAppOtaVersionsResponseBody struct {
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAppOtaVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppOtaVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppOtaVersionsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAppOtaVersionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAppOtaVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppOtaVersionsResponseBody) SetCode(v string) *DeleteAppOtaVersionsResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAppOtaVersionsResponseBody) SetMessage(v string) *DeleteAppOtaVersionsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAppOtaVersionsResponseBody) SetRequestId(v string) *DeleteAppOtaVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppOtaVersionsResponseBody) Validate() error {
	return dara.Validate(s)
}
