// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type EntityResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s EntityResponse) String() string {
	return tea.Prettify(s)
}

func (s EntityResponse) GoString() string {
	return s.String()
}

func (s *EntityResponse) SetHeaders(v map[string]*string) *EntityResponse {
	s.Headers = v
	return s
}

type IEResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s IEResponse) String() string {
	return tea.Prettify(s)
}

func (s IEResponse) GoString() string {
	return s.String()
}

func (s *IEResponse) SetHeaders(v map[string]*string) *IEResponse {
	s.Headers = v
	return s
}

type KWEResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s KWEResponse) String() string {
	return tea.Prettify(s)
}

func (s KWEResponse) GoString() string {
	return s.String()
}

func (s *KWEResponse) SetHeaders(v map[string]*string) *KWEResponse {
	s.Headers = v
	return s
}

type ReviewAnalysisResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s ReviewAnalysisResponse) String() string {
	return tea.Prettify(s)
}

func (s ReviewAnalysisResponse) GoString() string {
	return s.String()
}

func (s *ReviewAnalysisResponse) SetHeaders(v map[string]*string) *ReviewAnalysisResponse {
	s.Headers = v
	return s
}

type SentimentResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s SentimentResponse) String() string {
	return tea.Prettify(s)
}

func (s SentimentResponse) GoString() string {
	return s.String()
}

func (s *SentimentResponse) SetHeaders(v map[string]*string) *SentimentResponse {
	s.Headers = v
	return s
}

type TextStructureResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s TextStructureResponse) String() string {
	return tea.Prettify(s)
}

func (s TextStructureResponse) GoString() string {
	return s.String()
}

func (s *TextStructureResponse) SetHeaders(v map[string]*string) *TextStructureResponse {
	s.Headers = v
	return s
}

type TranslateResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s TranslateResponse) String() string {
	return tea.Prettify(s)
}

func (s TranslateResponse) GoString() string {
	return s.String()
}

func (s *TranslateResponse) SetHeaders(v map[string]*string) *TranslateResponse {
	s.Headers = v
	return s
}

type WordPosResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s WordPosResponse) String() string {
	return tea.Prettify(s)
}

func (s WordPosResponse) GoString() string {
	return s.String()
}

func (s *WordPosResponse) SetHeaders(v map[string]*string) *WordPosResponse {
	s.Headers = v
	return s
}

type WordSegmentResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
}

func (s WordSegmentResponse) String() string {
	return tea.Prettify(s)
}

func (s WordSegmentResponse) GoString() string {
	return s.String()
}

func (s *WordSegmentResponse) SetHeaders(v map[string]*string) *WordSegmentResponse {
	s.Headers = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("nlp"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) Entity(Domain *string) (_result *EntityResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &EntityResponse{}
	_body, _err := client.EntityWithOptions(Domain, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EntityWithOptions(Domain *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *EntityResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &EntityResponse{}
	_body, _err := client.DoROARequest(tea.String("Entity"), tea.String("2018-04-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/nlp/api/entity/"+tea.StringValue(Domain)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) IE(Domain *string) (_result *IEResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &IEResponse{}
	_body, _err := client.IEWithOptions(Domain, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) IEWithOptions(Domain *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *IEResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &IEResponse{}
	_body, _err := client.DoROARequest(tea.String("IE"), tea.String("2018-04-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/nlp/api/ie/"+tea.StringValue(Domain)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) KWE(Domain *string) (_result *KWEResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &KWEResponse{}
	_body, _err := client.KWEWithOptions(Domain, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) KWEWithOptions(Domain *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *KWEResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &KWEResponse{}
	_body, _err := client.DoROARequest(tea.String("KWE"), tea.String("2018-04-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/nlp/api/kwe/"+tea.StringValue(Domain)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReviewAnalysis(Domain *string) (_result *ReviewAnalysisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReviewAnalysisResponse{}
	_body, _err := client.ReviewAnalysisWithOptions(Domain, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReviewAnalysisWithOptions(Domain *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ReviewAnalysisResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &ReviewAnalysisResponse{}
	_body, _err := client.DoROARequest(tea.String("ReviewAnalysis"), tea.String("2018-04-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/nlp/api/reviewanalysis/"+tea.StringValue(Domain)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Sentiment(Domain *string) (_result *SentimentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SentimentResponse{}
	_body, _err := client.SentimentWithOptions(Domain, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SentimentWithOptions(Domain *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SentimentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &SentimentResponse{}
	_body, _err := client.DoROARequest(tea.String("Sentiment"), tea.String("2018-04-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/nlp/api/sentiment/"+tea.StringValue(Domain)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TextStructure(Domain *string) (_result *TextStructureResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TextStructureResponse{}
	_body, _err := client.TextStructureWithOptions(Domain, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TextStructureWithOptions(Domain *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TextStructureResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &TextStructureResponse{}
	_body, _err := client.DoROARequest(tea.String("TextStructure"), tea.String("2018-04-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/nlp/api/textstructure/"+tea.StringValue(Domain)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) Translate(Domain *string) (_result *TranslateResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TranslateResponse{}
	_body, _err := client.TranslateWithOptions(Domain, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TranslateWithOptions(Domain *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TranslateResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &TranslateResponse{}
	_body, _err := client.DoROARequest(tea.String("Translate"), tea.String("2018-04-08"), tea.String("HTTP"), tea.String("POST"), tea.String("AK"), tea.String("/nlp/api/translate/"+tea.StringValue(Domain)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) WordPos(Domain *string) (_result *WordPosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &WordPosResponse{}
	_body, _err := client.WordPosWithOptions(Domain, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) WordPosWithOptions(Domain *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *WordPosResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &WordPosResponse{}
	_body, _err := client.DoROARequest(tea.String("WordPos"), tea.String("2018-04-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/nlp/api/wordpos/"+tea.StringValue(Domain)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) WordSegment(Domain *string) (_result *WordSegmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &WordSegmentResponse{}
	_body, _err := client.WordSegmentWithOptions(Domain, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) WordSegmentWithOptions(Domain *string, headers map[string]*string, runtime *util.RuntimeOptions) (_result *WordSegmentResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	_result = &WordSegmentResponse{}
	_body, _err := client.DoROARequest(tea.String("WordSegment"), tea.String("2018-04-08"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("/nlp/api/wordsegment/"+tea.StringValue(Domain)), tea.String("none"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
