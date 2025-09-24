// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDatasetFileMetasStat interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DatasetFileMetasStat
	GetCount() *int32
	SetKey(v string) *DatasetFileMetasStat
	GetKey() *string
}

type DatasetFileMetasStat struct {
	// example:
	//
	// 10
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// cat
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
}

func (s DatasetFileMetasStat) String() string {
	return dara.Prettify(s)
}

func (s DatasetFileMetasStat) GoString() string {
	return s.String()
}

func (s *DatasetFileMetasStat) GetCount() *int32 {
	return s.Count
}

func (s *DatasetFileMetasStat) GetKey() *string {
	return s.Key
}

func (s *DatasetFileMetasStat) SetCount(v int32) *DatasetFileMetasStat {
	s.Count = &v
	return s
}

func (s *DatasetFileMetasStat) SetKey(v string) *DatasetFileMetasStat {
	s.Key = &v
	return s
}

func (s *DatasetFileMetasStat) Validate() error {
	return dara.Validate(s)
}
