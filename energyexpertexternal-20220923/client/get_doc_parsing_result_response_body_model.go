// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocParsingResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetDocParsingResultResponseBodyData) *GetDocParsingResultResponseBody
	GetData() *GetDocParsingResultResponseBodyData
	SetRequestId(v string) *GetDocParsingResultResponseBody
	GetRequestId() *string
}

type GetDocParsingResultResponseBody struct {
	// Returned result.
	Data *GetDocParsingResultResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// ID of the request
	//
	// example:
	//
	// 83A5A7DD-8974-5769-952E-590A97BEA34E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetDocParsingResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDocParsingResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetDocParsingResultResponseBody) GetData() *GetDocParsingResultResponseBodyData {
	return s.Data
}

func (s *GetDocParsingResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDocParsingResultResponseBody) SetData(v *GetDocParsingResultResponseBodyData) *GetDocParsingResultResponseBody {
	s.Data = v
	return s
}

func (s *GetDocParsingResultResponseBody) SetRequestId(v string) *GetDocParsingResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDocParsingResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDocParsingResultResponseBodyData struct {
	// - The parsed content of the document.
	//
	// - The format (markdown or json) is determined by the returnFormat parameter. For specific format details, refer to: [json return structure](https://www.alibabacloud.com/help/en/energy-expert/developer-reference/interface-attached-information#b644b6255cojj)
	//
	// example:
	//
	// {\\"doc_info\\":{\\"languages\\":[\\"zh\\",\\"en\\"],\\"doc_type\\":\\"pdf\\",\\"pdf_toc\\":[{\\"title\\":\\"封面\\",\\"level\\":0,\\"page\\":0}],\\"pages\\":366,\\"page_list\\":[{\\"imageWidth\\":596,\\"imageHeight\\":842,\\"pageIdAllDocs\\":0,\\"fileIndex\\":0,\\"pageIdCurDoc\\":0,\\"angle\\":0}],\\"doc_data\\":[{\\"uniqueId\\":\\"about_us_para\\",\\"page_num\\":\\"01\\",\\"index\\":\\"xxx\\",\\"name\\":\\"xxx\\",\\"type\\":\\"xxxx\\",\\"subType\\":\\"xxx\\",\\"text\\":\\"xxx\\",\\"before_text\\":\\"xxx\\",\\"after_text\\":\\"xxx\\",\\"extInfo\\":[{\\"uniqueId\\":\\"b0x1x0\\",\\"pos\\":[{\\"x\\":229,\\"y\\":208},{\\"x\\":421,\\"y\\":208},{\\"x\\":421,\\"y\\":242},{\\"x\\":229,\\"y\\":242}],\\"text\\":\\"Kurt Götze\\",\\"type\\":\\"Text\\",\\"subType\\":\\"Text\\",\\"pageNum\\":[0],\\"index\\":0}]}]}}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
}

func (s GetDocParsingResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDocParsingResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDocParsingResultResponseBodyData) GetContent() *string {
	return s.Content
}

func (s *GetDocParsingResultResponseBodyData) SetContent(v string) *GetDocParsingResultResponseBodyData {
	s.Content = &v
	return s
}

func (s *GetDocParsingResultResponseBodyData) Validate() error {
	return dara.Validate(s)
}
