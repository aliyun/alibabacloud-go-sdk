// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIndividuationTextResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDesc(v string) *DeleteIndividuationTextResponseBody
	GetDesc() *string
	SetRequestId(v string) *DeleteIndividuationTextResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteIndividuationTextResponseBody
	GetStatus() *string
}

type DeleteIndividuationTextResponseBody struct {
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4830493A-728F-5F19-BBCC-1443292E9C49
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DeleteIndividuationTextResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIndividuationTextResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIndividuationTextResponseBody) GetDesc() *string {
	return s.Desc
}

func (s *DeleteIndividuationTextResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIndividuationTextResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteIndividuationTextResponseBody) SetDesc(v string) *DeleteIndividuationTextResponseBody {
	s.Desc = &v
	return s
}

func (s *DeleteIndividuationTextResponseBody) SetRequestId(v string) *DeleteIndividuationTextResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIndividuationTextResponseBody) SetStatus(v string) *DeleteIndividuationTextResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteIndividuationTextResponseBody) Validate() error {
	return dara.Validate(s)
}
