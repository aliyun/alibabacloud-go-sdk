// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyModelApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModelApiId(v string) *ModifyModelApiResponseBody
	GetModelApiId() *string
	SetRequestId(v string) *ModifyModelApiResponseBody
	GetRequestId() *string
}

type ModifyModelApiResponseBody struct {
	// example:
	//
	// mi-xxxxx
	ModelApiId *string `json:"ModelApiId,omitempty" xml:"ModelApiId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 925B84D9-CA72-432C-95CF-738C22******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyModelApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyModelApiResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyModelApiResponseBody) GetModelApiId() *string {
	return s.ModelApiId
}

func (s *ModifyModelApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyModelApiResponseBody) SetModelApiId(v string) *ModifyModelApiResponseBody {
	s.ModelApiId = &v
	return s
}

func (s *ModifyModelApiResponseBody) SetRequestId(v string) *ModifyModelApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyModelApiResponseBody) Validate() error {
	return dara.Validate(s)
}
