// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostInnerReaderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataSourceDescriptor(v *PostInnerReaderRequestDataSourceDescriptor) *PostInnerReaderRequest
	GetDataSourceDescriptor() *PostInnerReaderRequestDataSourceDescriptor
	SetDataSourceName(v string) *PostInnerReaderRequest
	GetDataSourceName() *string
}

type PostInnerReaderRequest struct {
	DataSourceDescriptor *PostInnerReaderRequestDataSourceDescriptor `json:"dataSourceDescriptor,omitempty" xml:"dataSourceDescriptor,omitempty" type:"Struct"`
	// example:
	//
	// ds_dolphin_prod
	DataSourceName *string `json:"dataSourceName,omitempty" xml:"dataSourceName,omitempty"`
}

func (s PostInnerReaderRequest) String() string {
	return dara.Prettify(s)
}

func (s PostInnerReaderRequest) GoString() string {
	return s.String()
}

func (s *PostInnerReaderRequest) GetDataSourceDescriptor() *PostInnerReaderRequestDataSourceDescriptor {
	return s.DataSourceDescriptor
}

func (s *PostInnerReaderRequest) GetDataSourceName() *string {
	return s.DataSourceName
}

func (s *PostInnerReaderRequest) SetDataSourceDescriptor(v *PostInnerReaderRequestDataSourceDescriptor) *PostInnerReaderRequest {
	s.DataSourceDescriptor = v
	return s
}

func (s *PostInnerReaderRequest) SetDataSourceName(v string) *PostInnerReaderRequest {
	s.DataSourceName = &v
	return s
}

func (s *PostInnerReaderRequest) Validate() error {
	if s.DataSourceDescriptor != nil {
		if err := s.DataSourceDescriptor.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PostInnerReaderRequestDataSourceDescriptor struct {
	// example:
	//
	// test_ds318_hangzhou_0428
	DsName *string `json:"dsName,omitempty" xml:"dsName,omitempty"`
}

func (s PostInnerReaderRequestDataSourceDescriptor) String() string {
	return dara.Prettify(s)
}

func (s PostInnerReaderRequestDataSourceDescriptor) GoString() string {
	return s.String()
}

func (s *PostInnerReaderRequestDataSourceDescriptor) GetDsName() *string {
	return s.DsName
}

func (s *PostInnerReaderRequestDataSourceDescriptor) SetDsName(v string) *PostInnerReaderRequestDataSourceDescriptor {
	s.DsName = &v
	return s
}

func (s *PostInnerReaderRequestDataSourceDescriptor) Validate() error {
	return dara.Validate(s)
}
