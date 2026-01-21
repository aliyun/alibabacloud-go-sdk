// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBsnByResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatas(v *GetBsnByResourceResponseBodyDatas) *GetBsnByResourceResponseBody
	GetDatas() *GetBsnByResourceResponseBodyDatas
}

type GetBsnByResourceResponseBody struct {
	Datas *GetBsnByResourceResponseBodyDatas `json:"datas,omitempty" xml:"datas,omitempty" type:"Struct"`
}

func (s GetBsnByResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBsnByResourceResponseBody) GoString() string {
	return s.String()
}

func (s *GetBsnByResourceResponseBody) GetDatas() *GetBsnByResourceResponseBodyDatas {
	return s.Datas
}

func (s *GetBsnByResourceResponseBody) SetDatas(v *GetBsnByResourceResponseBodyDatas) *GetBsnByResourceResponseBody {
	s.Datas = v
	return s
}

func (s *GetBsnByResourceResponseBody) Validate() error {
	if s.Datas != nil {
		if err := s.Datas.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetBsnByResourceResponseBodyDatas struct {
	BsnDO []*string `json:"bsnDO,omitempty" xml:"bsnDO,omitempty" type:"Repeated"`
}

func (s GetBsnByResourceResponseBodyDatas) String() string {
	return dara.Prettify(s)
}

func (s GetBsnByResourceResponseBodyDatas) GoString() string {
	return s.String()
}

func (s *GetBsnByResourceResponseBodyDatas) GetBsnDO() []*string {
	return s.BsnDO
}

func (s *GetBsnByResourceResponseBodyDatas) SetBsnDO(v []*string) *GetBsnByResourceResponseBodyDatas {
	s.BsnDO = v
	return s
}

func (s *GetBsnByResourceResponseBodyDatas) Validate() error {
	return dara.Validate(s)
}
