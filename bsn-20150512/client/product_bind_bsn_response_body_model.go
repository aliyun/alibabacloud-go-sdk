// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProductBindBsnResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatas(v *ProductBindBsnResponseBodyDatas) *ProductBindBsnResponseBody
	GetDatas() *ProductBindBsnResponseBodyDatas
	SetRequestId(v string) *ProductBindBsnResponseBody
	GetRequestId() *string
}

type ProductBindBsnResponseBody struct {
	Datas *ProductBindBsnResponseBodyDatas `json:"datas,omitempty" xml:"datas,omitempty" type:"Struct"`
	// example:
	//
	// 21ED07AF-267E-5820-AEE5-53C973BFD5F8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ProductBindBsnResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ProductBindBsnResponseBody) GoString() string {
	return s.String()
}

func (s *ProductBindBsnResponseBody) GetDatas() *ProductBindBsnResponseBodyDatas {
	return s.Datas
}

func (s *ProductBindBsnResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ProductBindBsnResponseBody) SetDatas(v *ProductBindBsnResponseBodyDatas) *ProductBindBsnResponseBody {
	s.Datas = v
	return s
}

func (s *ProductBindBsnResponseBody) SetRequestId(v string) *ProductBindBsnResponseBody {
	s.RequestId = &v
	return s
}

func (s *ProductBindBsnResponseBody) Validate() error {
	if s.Datas != nil {
		if err := s.Datas.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ProductBindBsnResponseBodyDatas struct {
	BsnDO []*string `json:"bsnDO,omitempty" xml:"bsnDO,omitempty" type:"Repeated"`
}

func (s ProductBindBsnResponseBodyDatas) String() string {
	return dara.Prettify(s)
}

func (s ProductBindBsnResponseBodyDatas) GoString() string {
	return s.String()
}

func (s *ProductBindBsnResponseBodyDatas) GetBsnDO() []*string {
	return s.BsnDO
}

func (s *ProductBindBsnResponseBodyDatas) SetBsnDO(v []*string) *ProductBindBsnResponseBodyDatas {
	s.BsnDO = v
	return s
}

func (s *ProductBindBsnResponseBodyDatas) Validate() error {
	return dara.Validate(s)
}
